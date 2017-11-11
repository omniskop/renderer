// package rieger859663
package main

import (
    "log"
    "fmt"
    "math"
    "sync"
    "runtime"
    "time"
    "cgtools/random"
    "cgtools/image"
    "customtools/vec3"
    "customtools/shapes"
    "customtools/camera"
    "customtools/ray"
)

var startTime time.Time;

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    startTime = time.Now()
    random.InitSeed();
    
    world := shapes.Group{
        []shapes.Shape{
            shapes.Sphere{
                Position: vec3.Vec3{0,-0.1,-11.2},
                Material: shapes.Material_Diffuse{vec3.Red},
                Radius: 1,
            },
            shapes.Sphere{
                Position: vec3.Vec3{-.7,-.7,-10.7},
                Material: shapes.Material_Diffuse{vec3.Blue},
                Radius: .5,
            },
            shapes.Sphere{
                Position: vec3.Vec3{.7,-.7,-10.7},
                Material: shapes.Material_Diffuse{vec3.Blue},
                Radius: .5,
            },
            shapes.Plane{
                Position: vec3.Vec3{0,-1,0},
                Normal: vec3.Vec3{0,1,0},
                Material: shapes.Material_Diffuse{vec3.Grey},
            },
            shapes.Background{shapes.Material_Sky{vec3.White}},
        },
    }
    
    img := raytrace(
        camera.PinholeCamera{
            OpeningAngle: math.Pi / 13,
            Width: 200,
            Height: 200,
        },
        world,
        200,
        4,
    )
    log.Print("Rendering took ", time.Since(startTime))
        
    err := img.Write("doc/a05-diffuse-spheres.png")
    if err != nil {
        log.Print("An error occoured while writing the file.")
        log.Fatal( err )
    } else {
        log.Print("File saved.")
    }
}

func raytrace(cam camera.PinholeCamera, shapes shapes.Group, sPoints, depth int) image.Image {
    img := image.New(cam.Width, cam.Height)
    pointPerPixelAxis := int(math.Sqrt(float64(sPoints)))
    var wg sync.WaitGroup
    dataChannel := make(chan PixelInformation, cam.Width * cam.Height)
    
    wg.Add(cam.Width * cam.Height)
    for x := 0; x < cam.Width;x++ {
        for y := 0; y < cam.Height;y++ {
            // img.SetPixel(
            //     x, 
            //     y,
            //     processColor(
            //         getColorForPixel(shapes, cam, x, y, pointPerPixelAxis, depth),
            //         2.2,
            //     ),
            // )
            go func(px,py int){
                defer wg.Done()
            //     // img.SetPixel(
            //     //     px, 
            //     //     py,
            //     //     processColor(
            //     //         getColorForPixel(shapes, cam, x, y, pointPerPixelAxis),
            //     //         2.2,
            //     //     ),
            //     // )
                dataChannel <- PixelInformation{
                    px,
                    py,
                    processColor(
                        getColorForPixel(shapes, cam, px, py, pointPerPixelAxis, depth),
                        2.2,
                    ),
                }
            }(x,y)
        }
        // percent := (float64(x) / float64(cam.Width)) * 100
        // if percent == math.Trunc(percent) {
        //     fmt.Printf("\rRendering... %d%%", int(percent))
        // }
    }
    
    go func() {
        i := 0
        for val := range dataChannel {
            img.SetPixel(
                val.X,
                val.Y,
                val.Color,
            )
            i++
            percent := (float64(i) / float64(cam.Width * cam.Height)) * 100
            if percent == math.Trunc(percent) {
                fmt.Printf("\rRendering... %d%%", int(percent))
            }
        }
    }()
    
    wg.Wait();

    
    fmt.Printf("\n")

    return img
}

func SetPixel(x,y int ,waitGroup *sync.WaitGroup,dataChannel *chan PixelInformation, shapes shapes.Shape, cam camera.PinholeCamera, pointPerPixelAxis, depth int)  {
    // img.SetPixel(
    //     x, 
    //     y,
    //     processColor(
    //         getColorForPixel(shapes, cam, x, y, pointPerPixelAxis),
    //         2.2,
    //     ),
    // )
    defer waitGroup.Done()
    *dataChannel <- PixelInformation{
        x,
        y,
        processColor(
            getColorForPixel(shapes, cam, x, y, pointPerPixelAxis, depth),
            2.2,
        ),
    }
}

func getColorForPixel(shapes shapes.Shape, cam camera.PinholeCamera,pX, pY, supersamplingPoints, depth int) vec3.Vec3 {
    color := vec3.Vec3{0,0,0}
    for x := 0; x < supersamplingPoints;x++ {
        for y := 0; y < supersamplingPoints;y++ {
            pixelRay := cam.GetRayForPixel(
                // float64(pX) + float64(x) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
                // float64(pY) + float64(y) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
                float64(pX) + float64(x) / float64(supersamplingPoints) + random.Float64() * 0 + .5 / float64(supersamplingPoints),
                float64(pY) + float64(y) / float64(supersamplingPoints) + random.Float64() * 0 + .5 / float64(supersamplingPoints),
            )
            /*closestHit := shapes.Intersect(pixelRay)
            
            if closestHit == nil {
                color.Add(vec3.Black)
                log.Print("Warning: No background present.")
            } else {
                color.Add(closestHit.Material.Albedo(pixelRay, *closestHit));
                // color.Add(vec3.Red)
            }*/
            
            c := calculateRadiance(shapes, pixelRay, depth)
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

func processColor(color vec3.Vec3, gamma float64) vec3.Vec3 {
    return vec3.Vec3{
        math.Pow( color.X, 1 / gamma ),
        math.Pow( color.Y, 1 / gamma ),
        math.Pow( color.Z, 1 / gamma ),
    };
}

type PixelInformation struct {
    X   int
    Y   int
    Color   vec3.Vec3
}




