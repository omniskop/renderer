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

const width int = 1000
const height int = 1000
const supersamplingPoints = 10
var backgroundColor = vec3.Black
// var spheres = make([]circle.Circle, 20);
var globe shapes.Sphere;
var cam = camera.PinholeCamera{
    OpeningAngle: math.Pi / 16,
    Width: width,
    Height: height,
}

func main() {
    // random.InitSeed()
    // for i, _ := range circles {
    //     spheres[i] = Sphere.NewRandom(width, height)
    // }
    // spheres.SortByRadius(circles)
    
    globe = shapes.Sphere{
        Position: vec3.Vec3{0,0,-11},
        Radius: 1,
    }
    
    img := image.New(width, height)
    
    for x := 0; x < width;x++ {
        for y := 0; y < height;y++ {
            img.SetPixel(x, y, pixelColor(x, y))
        }
    }
        
    err := img.Write("doc/a03/a03-one-sphere.png")
    if err != nil {
        log.Print("An error occoured while writing the file.")
        log.Fatal( err )
    } else {
        log.Print("File saved.")
    }
}

func pixelColor(pX int, pY int) vec3.Vec3 {
    color := vec3.Vec3{0,0,0}
    for x := 0; x < supersamplingPoints;x++ {
        for y := 0; y < supersamplingPoints;y++ {
            pixelRay := cam.GetRayForPixel(
                float64(pX) + float64(x) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
                float64(pY) + float64(y) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
            )
            clostestHit := globe.Intersect(pixelRay)
            
            if clostestHit == nil {
                color.Add(backgroundColor)
            } else {
                color.Add(clostestHit.Normal)
            }
        }
    }

    color.Divide( supersamplingPoints * supersamplingPoints )
    
    return color
}





