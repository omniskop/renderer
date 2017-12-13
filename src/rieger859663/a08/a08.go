// package rieger859663
package main

import (
    "log"
    "fmt"
    "math"
    "time"
    "runtime"
    "sync"
    "flag"
    "os"
    "runtime/pprof"
    "cgtools/random"
    "cgtools/image"
    "customtools/vec3"
    "customtools/shapes"
    "customtools/camera"
    "customtools/ray"
    "customtools/space"
    "scenes"
)

var startTime time.Time;

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")


func main() {
    flag.Parse()
    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()
    }
    startTime = time.Now()
    random.InitSeed();
    
    _ = scenes.GetWaterMolecule()
    _ = space.NewTransformation
    
    
    // sceneCamera := scenes.GetDnaSceneCamera()
    // sceneCamera := scenes.GetCylinderSceneCamera()
    sceneCamera := scenes.GetHumanDnaCamera()
    
    //Standart Scenen
    img := multithreadMagic(
        camera.PinholeCamera{
            Position: sceneCamera.GetPosition(),
            Direction: sceneCamera.GetDirection(),
            OpeningAngle: sceneCamera.GetOpeningAngle(),
            Width: 500,
            Height: 500,
        },
        scenes.GetHumanDnaScene(),
        100,
        20,
    )
    
    // Kamera tests
    // img := multithreadMagic(
    //     camera.PinholeCamera{
    //         Position: vec3.Vec3{2,1,-10},
    //         Direction: vec3.Normalize(vec3.Vec3{-0.4,-0.3,-1}),
    //         Tilt: 0,
    //         OpeningAngle: math.Pi / 7,
    //         Width: 200,
    //         Height: 200,
    //     },
    //     scenes.GetComparisonScene(),
    //     100,
    //     30,
    // )

    // Performance tests
    // img := raytrace(
    //     camera.PinholeCamera{
    //         Position: sceneCamera.GetPosition(),
    //         Direction: sceneCamera.GetDirection(),
    //         OpeningAngle: sceneCamera.GetOpeningAngle(),
    //         Width: 200,
    //         Height: 200,
    //     },
    //     scenes.GetCylinderScene(),
    //     300,
    //     30,
    // )
    
    //scene tests
    /*img := multithreadMagic(
        camera.PinholeCamera{
                Position: vec3.Vec3{0,4,5},
                Direction: vec3.Normalize(vec3.Vec3{0,-0.8,-1}),
                OpeningAngle: math.Pi / 2,
                Width: 300,
                Height: 300,
            },
            shapes.Group{
                // space.NewTransformation(vec3.Vec3{0,0,0}, math.Pi * 0,0,0),
                space.NoTransformation(),
                []shapes.Shape{
                    shapes.Group{
                        space.NewTransformation(vec3.Vec3{0,0,0}, deg2rad(0),deg2rad(0),deg2rad(45)),
                        // space.NoTransformation(),
                        []shapes.Shape{
                            // shapes.NewCylinder(
                            //     vec3.Vec3{0,1,0},
                            //     vec3.Vec3{0,1,0},
                            //     1,
                            //     5,
                            //     space.Material_Normal{},
                            // ),
                            shapes.NewCone(
                                vec3.Vec3{0,1,0},
                                vec3.Vec3{0,1,0},
                                1,
                                3,
                                space.Material_Diffuse{vec3.Red},
                                // space.Material_Normal{},
                            ),
                            &shapes.Sphere{
                                vec3.Vec3{0,0,0},
                                1,
                                space.Material_Diffuse{vec3.Green},
                                // space.Material_Normal{},
                            },
                        },
                    },
                    // shapes.NewCylinder(
                    //     vec3.Vec3{1,0,-2},
                    //     vec3.Vec3{1,0,0},
                    //     1,
                    //     5,
                    //     space.Material_Normal{},
                    // ),
                    &shapes.Sphere{
                        vec3.Vec3{0,0,0},
                        0.5,
                        space.Material_Diffuse{vec3.White},
                    },
                    &shapes.Plane{
                        vec3.Vec3{0,0,0},
                        vec3.Vec3{0,1,0},
                        space.Material_Diffuse{vec3.Blue},
                    },
                    shapes.Background{space.Material_Sky{vec3.White}},
                },
            },
            40,
            10,
    )*/
    
    log.Print("Rendering took ", time.Since(startTime))
        
    err := img.Write("doc/a08-1.png")
    if err != nil {
        log.Print("An error occoured while writing the file.")
        log.Fatal( err )
    } else {
        log.Print("File saved.")
    }
}

func deg2rad(x float64) float64 {
    return (x/360) * math.Pi*2
}

func multithreadMagic(cam camera.Camera, scene shapes.Shape, supersamplingPoints , depth int) image.Image {
    // threads := runtime.NumCPU()
    threads := 7
    runtime.GOMAXPROCS(threads)
    log.Print(threads, " threads")
    
    samplingPointsPerThread := supersamplingPoints / threads
    log.Print(samplingPointsPerThread * threads, " sampling points")
    log.Print(samplingPointsPerThread, " sampling points per thread")
    
    images := make([]image.Image,threads)
    
    var wg sync.WaitGroup
    channel := make(chan image.Image)
    
    for i := 0;i < threads;i++ {
        wg.Add(1)
        go renderThread(&wg,channel, cam, scene, samplingPointsPerThread, depth)
    }
    
    for i := 0;i < threads;i++ {
        images[i] = <- channel
    }
    
    wg.Wait()
    
    finalImage := image.New(cam.GetWidth(), cam.GetHeight())
    
    threadsFloat := float64(threads)
    for x:= 0;x < cam.GetWidth();x++ {
        for y := 0; y < cam.GetHeight();y++ {
            color := vec3.Black
            for _, img := range images {
                color.Add(img.GetPixel(x,y))
            }
            finalImage.SetPixel(x,y, processColor(vec3.Divide(color, threadsFloat),2.2) )
        }
    }    
    
    return finalImage
}

func renderThread(wg *sync.WaitGroup, out chan image.Image,cam camera.Camera, scene shapes.Shape, sPoints, depth int) {
    img := raytrace(cam, scene, sPoints, depth)
    out <- img
    wg.Done()
}

func raytrace(cam camera.Camera, shapes shapes.Shape, sPoints, depth int) image.Image {
    img := image.New(cam.GetWidth(), cam.GetHeight())
    pointPerPixelAxis := int(math.Sqrt(float64(sPoints)))
    for x := 0; x < cam.GetWidth();x++ {
        for y := 0; y < cam.GetHeight();y++ {
            img.SetPixel(
                x, 
                y,
                // processColor(
                    getColorForPixel(shapes, cam, x, y, pointPerPixelAxis, depth),
                    // 2.2,
                // ),

            )
        }
        percent := (float64(x) / float64(cam.GetWidth())) * 100
        if percent == math.Trunc(percent) {
            fmt.Printf("\rRendering... %d%%", int(percent))
        }
    }

    
    fmt.Printf("\n")

    return img
}

func getColorForPixel(shapes shapes.Shape, cam camera.Camera,pX, pY, supersamplingPoints, depth int) vec3.Vec3 {
    color := vec3.Vec3{0,0,0}
    for x := 0; x < supersamplingPoints;x++ {
        for y := 0; y < supersamplingPoints;y++ {
            pixelRay := cam.GetRayForPixel(
                float64(pX) + (float64(x) + random.Float64()) / float64(supersamplingPoints),
                float64(pY) + (float64(y) + random.Float64()) / float64(supersamplingPoints),
            )
            
            // c := calculateRadiance(shapes, pixelRay, depth)
            c := calculateRadiance2(shapes, pixelRay, depth)
            color.Add(c)
        }
    }

    color.Divide( float64(supersamplingPoints * supersamplingPoints) )
    
    return color
}

func calculateRadiance(scene shapes.Shape, renderRay ray.Ray, depth int) vec3.Vec3 {
    closestHit := scene.Intersect(renderRay);
    
    radiance := vec3.Zero;
    
    if closestHit == nil {
        log.Print("Sollte nicht passieren.")
        return radiance;
    }
    
    emission := closestHit.Material.EmittedRadiance(renderRay, *closestHit);
    scattered := closestHit.Material.ScatteredRay(renderRay, *closestHit);
    
    if scattered != nil {
        depth--;
        if depth != 0 {
            radiance = calculateRadiance(scene, *scattered, depth);
        }
        return vec3.Add(
            emission,
            vec3.MultiplyByVec3(
                closestHit.Material.Albedo(renderRay, *closestHit),
                radiance,
            ),
        )
    } else {
        return emission;
    }
}

func calculateRadiance2(scene shapes.Shape, renderRay ray.Ray, depth int) vec3.Vec3 {
    currentRay := &renderRay
    value := vec3.Zero
    add := vec3.One
    for i := 0; i < depth;i++ {
        closestHit := scene.Intersect(*currentRay);
        if closestHit == nil {
            log.Print("Sollte nicht passieren.")
            break
        }
        
        emission := closestHit.Material.EmittedRadiance(*currentRay, *closestHit);
        // c := closestHit.Position.Y / 1
        // if c < 0 {
        //     c = 0
        // }
        // emission := vec3.Vec3{c,c,c}
        currentRay = closestHit.Material.ScatteredRay(*currentRay, *closestHit);
        
        value.Add(vec3.MultiplyByVec3(add, emission))
        if currentRay != nil {
            add.MultiplyByVec3(closestHit.Material.Albedo(renderRay, *closestHit))
        } else {
            break
        }
    }
    return value
}

func processColor(color vec3.Vec3, gamma float64) vec3.Vec3 {
    return vec3.Vec3{
        math.Pow( color.X, 1 / gamma ),
        math.Pow( color.Y, 1 / gamma ),
        math.Pow( color.Z, 1 / gamma ),
    };
}




