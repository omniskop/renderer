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
var backgroundColor = vec3.Vec3{1,1,1}
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
        
    err := img.Write("doc/a02/a02-disks.png")
    if err != nil {
        log.Print("An error occoured while writing the file.")
        log.Fatal( err )
    } else {
        log.Print("File saved.")
    }
}

func pixelColor(x int, y int) vec3.Vec3 {
    for _, c := range circles {
        if distance(x, y, c.Pos.X, c.Pos.Y) < c.Radius {
            return c.Color
        }
    }
    return backgroundColor
}

func distance(ax, ay, bx, by int) int {
    return int( math.Sqrt( math.Pow( float64(ax - bx), 2) + math.Pow( float64(ay - by), 2) ) )
}




