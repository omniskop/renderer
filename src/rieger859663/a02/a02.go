// package rieger859663
package main

import (
    "log"
    "math"
    "cgtools/image"
    "cgtools/vec3"
    "cgtools/circle"
    "cgtools/random"
)

const width int = 160
const height int = 90
const supersamplingPoints = 10
var backgroundColor = vec3.White
var circles = make([]circle.Circle, 20);

func main() {    
    random.InitSeed()
    
    for i, _ := range circles {
        circles[i] = circle.NewRandom(width, height)
    }
    circle.SortByRadius(circles)
    
    img := image.New(width, height)
    
    for x := 0; x < width;x++ {
        for y := 0; y < height;y++ {
            img.SetPixel(x, y, pixelColor(x, y))
        }
    }
        
    err := img.Write("doc/a02/a02-super-sampling.png")
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
            for _, c := range circles {
                if checkSubpixel(pX, pY, x, y, c) {
                    color.Add(c.Color)
                    goto noBackground
                }
            }
            color.Add(backgroundColor)
            noBackground:
        }
    }

    color.Divide( supersamplingPoints * supersamplingPoints )

    return color
}

func checkSubpixel(pX, pY, x, y int, c circle.Circle) bool {
    return distance(
        float64(pX) + float64(x) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
        float64(pY) + float64(y) / float64(supersamplingPoints) + random.Float64() / float64(supersamplingPoints),
        float64(c.Pos.X),
        float64(c.Pos.Y),
    ) < float64(c.Radius)
}

func distance(ax, ay, bx, by float64) float64 {
    return math.Sqrt( math.Pow( ax - bx, 2) + math.Pow( ay - by, 2) )
}







