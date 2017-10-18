// package rieger859663
package main

import (
    "log"
    "cgtools/image"
    "cgtools/vec3"
)

const width int = 160
const height int = 90

func main() {    
    img := image.New(width, height)
    
    for x := 0; x < width;x++ {
        for y := 0; y < height;y++ {
            img.SetPixel(x, y, pixelColor(x, y))
        }
    }
    
    err := img.Write("doc/a01-checkerboard.png")
    if err != nil {
        log.Print("An error occoured while writing the file.")
        log.Fatal( err )
    } else {
        log.Print("File saved.")
    }
}

func pixelColor(x int, y int) vec3.Vec3 {
    if ((x / 10) % 2 == 0) != ((y / 10) % 2 == 0)  {
        return vec3.Vec3{0,0,1}
    } else {
        return vec3.Vec3{
            float64(x) / float64(width),
            1- float64( x) / float64(width),
            .5,
        }
    }
}