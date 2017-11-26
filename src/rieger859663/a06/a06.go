// package rieger859663
package main

import (
    "log"
    "fmt"
    "math"
    "time"
    "cgtools/random"
    "cgtools/image"
    "customtools/vec3"
    "customtools/shapes"
    "customtools/camera"
    "customtools/ray"
    _"scenes"
)

var startTime time.Time;

func main() {    
    startTime = time.Now()
    random.InitSeed();
    
    // world := scenes.GetWaterMolecule()
    
    world := shapes.Group {
        []shapes.Shape{
            shapes.NewCylinder(
                vec3.Vec3{0,-1.3,-5},
                vec3.Normalize(vec3.Vec3{0,1,0}),
                1,
                2,
                shapes.Material_Metal{vec3.White, 0},
            ),
            shapes.NewCylinder(
                vec3.Vec3{-1,-1.3,-8},
                vec3.Normalize(vec3.Vec3{-1,1,-1}),
                0.5,
                2,
                shapes.Material_Diffuse{vec3.Red},
            ),
            shapes.NewCylinder(
                vec3.Vec3{1,-1.3,-8},
                vec3.Normalize(vec3.Vec3{1,1,-1}),
                0.5,
                2,
                shapes.Material_Diffuse{vec3.Red},
            ),
            shapes.NewCylinder(
                vec3.Vec3{0,-1.3,-3},
                vec3.Normalize(vec3.Vec3{0,1,0}),
                0.5,
                0.3,
                shapes.Material_Diffuse{vec3.Vec3{0.25,0.75,1}},
            ),
            shapes.Plane{ // bottom
                Position: vec3.Vec3{0,-1.3,0},
                Normal: vec3.Vec3{0,1,0},
                Material: shapes.Material_Diffuse_Checkerboard{1,vec3.White, vec3.Black},
            },
            shapes.Background{shapes.Material_Sky{vec3.Vec3{0.8,0.8,1}}},
        },
    }
    
    img := raytrace(
        camera.PinholeCamera{
            Position: vec3.Vec3{0,2,7},
            Direction: vec3.Normalize(vec3.Vec3{0,-0.15,-1}),
            OpeningAngle: math.Pi / 9,
            Width: 500,
            Height: 500,
        },
        world,
        200,
        10,
    )
    log.Print("Rendering took ", time.Since(startTime))
        
    err := img.Write("doc/a06-disc.png")
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
    for x := 0; x < cam.Width;x++ {
        for y := 0; y < cam.Height;y++ {
            img.SetPixel(
                x, 
                y,
                processColor(
                    getColorForPixel(shapes, cam, x, y, pointPerPixelAxis, depth),
                    2.2,
                ),
            )
        }
        percent := (float64(x) / float64(cam.Width)) * 100
        if percent == math.Trunc(percent) {
            fmt.Printf("\rRendering... %d%%", int(percent))
        }
    }

    
    fmt.Printf("\n")

    return img
}

func getColorForPixel(shapes shapes.Shape, cam camera.PinholeCamera,pX, pY, supersamplingPoints, depth int) vec3.Vec3 {
    color := vec3.Vec3{0,0,0}
    for x := 0; x < supersamplingPoints;x++ {
        for y := 0; y < supersamplingPoints;y++ {
            pixelRay := cam.GetRayForPixel(
                float64(pX) + float64(x) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
                float64(pY) + float64(y) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
                // float64(pX) + float64(x) / float64(supersamplingPoints) + random.Float64() * 0 + .5 / float64(supersamplingPoints),
                // float64(pY) + float64(y) / float64(supersamplingPoints) + random.Float64() * 0 + .5 / float64(supersamplingPoints),
            )
            
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




