// package rieger859663
package main

import (
    "log"
    "math"
    "cgtools/image"
    "customtools/vec3"
    "customtools/sphere"
    "customtools/camera"
    "customtools/hit"
)

const width int = 1000
const height int = 1000
const supersamplingPoints = 10
var backgroundColor = vec3.Black
// var spheres = make([]circle.Circle, 20);
var spheres = make([]sphere.Sphere, 1);
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
    
    spheres[0] = sphere.Sphere{
        Position: vec3.Vec3{0,0,-11},
        Radius: 1,
    }
    
    img := image.New(width, height)
    
    for x := 0; x < width;x++ {
        for y := 0; y < height;y++ {
            img.SetPixel(x, y, pixelColor(x, y))
        }
    }
        
    err := img.Write("doc/a04-3-spheres.png")
    if err != nil {
        log.Print("An error occoured while writing the file.")
        log.Fatal( err )
    } else {
        log.Print("File saved.")
    }
}

func pixelColor(x int, y int) vec3.Vec3 {
    color := vec3.Vec3{0,0,0}
    
    pixelRay := cam.GetRayForPixel(x,y)
    var clostestHit *hit.Hit
    
    for _, currentSphere := range spheres {
        sphereHit := currentSphere.Intersect(pixelRay)
        if sphereHit != nil {
            if clostestHit == nil || sphereHit.Position.LessThan(clostestHit.Position) {
                clostestHit = sphereHit
            }
        }
    }
    
    if clostestHit == nil {
        color = backgroundColor
    } else {
        color = clostestHit.Normal
    }
    
    return color
}






