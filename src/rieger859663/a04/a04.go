// package rieger859663
package main

import (
    "log"
    "math"
    "cgtools/random"
    "cgtools/image"
    "customtools/vec3"
    "customtools/shapes"
    "customtools/camera"
)

func main() {
    
    world := shapes.Group{
        []shapes.Shape{
            shapes.Sphere{
                Position: vec3.Vec3{0,-0.1,-11.2},
                Color: vec3.Red,
                Radius: 1,
            },
            shapes.Sphere{
                Position: vec3.Vec3{-.7,-.7,-10.7},
                Color: vec3.Blue,
                Radius: .5,
            },
            shapes.Sphere{
                Position: vec3.Vec3{.7,-.7,-10.7},
                Color: vec3.Blue,
                Radius: .5,
            },
            shapes.Plane{
                Position: vec3.Vec3{0,-1,0},
                Normal: vec3.Vec3{0,1,0},
                Color: vec3.Vec3{.5,.5,.5},
            },
            shapes.Background{vec3.White},
        },
    }
    
    img := raytrace(
        camera.PinholeCamera{
            OpeningAngle: math.Pi / 13,
            Width: 1000,
            Height: 1000,
        },
        world,
        100,
    )
        
    // world := shapes.Group{
    //     []shapes.Shape{
    //         shapes.Sphere{
    //             Position: vec3.Vec3{-1.0,-0.25,-2.5},
    //             Color: vec3.Red,
    //             Radius: .7,
    //         },
    //         shapes.Sphere{
    //             Position: vec3.Vec3{0,-0.25,-2.5},
    //             Color: vec3.Green,
    //             Radius: .5,
    //         },
    //         shapes.Sphere{
    //             Position: vec3.Vec3{1,-0.25,-2.5},
    //             Color: vec3.Blue,
    //             Radius: .7,
    //         },
    //         shapes.Plane{
    //             Position: vec3.Vec3{0,-0.5,0},
    //             Normal: vec3.Vec3{0,1,0},
    //             Color: vec3.Grey,
    //         },
    //         shapes.Background{vec3.White},
    //     },
    // }
    // 
    // img := raytrace(
    //     camera.PinholeCamera{
    //         OpeningAngle: math.Pi / 2,
    //         Width: 2000,
    //         Height: 1500,
    //     },
    //     world,
    //     100,
    // )
        
    err := img.Write("doc/a04-scene.png")
    if err != nil {
        log.Print("An error occoured while writing the file.")
        log.Fatal( err )
    } else {
        log.Print("File saved.")
    }
}

func raytrace(cam camera.PinholeCamera, shapes shapes.Group, sPoints int) image.Image {
    img := image.New(cam.Width, cam.Height)
    pointPerPixelAxis := int(math.Sqrt(float64(sPoints)))
    
    for x := 0; x < cam.Width;x++ {
        for y := 0; y < cam.Height;y++ {
            img.SetPixel(
                x, 
                y,
                processColor(
                    getColorForPixel(shapes, cam, x, y, pointPerPixelAxis),
                    2.2,
                ),
            )
        }
    }
    
    return img
}

func getColorForPixel(shapes shapes.Shape, cam camera.PinholeCamera,pX, pY, supersamplingPoints int) vec3.Vec3 {
    color := vec3.Vec3{0,0,0}
    for x := 0; x < supersamplingPoints;x++ {
        for y := 0; y < supersamplingPoints;y++ {
            pixelRay := cam.GetRayForPixel(
                float64(pX) + float64(x) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
                float64(pY) + float64(y) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
            )
            clostestHit := shapes.Intersect(pixelRay)
            
            if clostestHit == nil {
                color.Add(vec3.Black)
                log.Print("Warning: No background present.")
            } else {
                color.Add(clostestHit.Color)
            }
        }
    }

    color.Divide( float64(supersamplingPoints * supersamplingPoints) )
    
    return color
}

func processColor(color vec3.Vec3, gamma float64) vec3.Vec3 {
    return vec3.Vec3{
        math.Pow( color.X, 1 / gamma ),
        math.Pow( color.Y, 1 / gamma ),
        math.Pow( color.Z, 1 / gamma ),
    };
}





