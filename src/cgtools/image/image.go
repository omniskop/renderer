package image

import (
    "cgtools/vec3"
    "cgtools/imageWriter"
    "math"
)

const gamma = 2.2

type Image struct {
    width int
    height int
    imageData []float64
}

func New(w int, h int) *Image {
    return &Image{
        width: w,
        height: h,
        imageData: make([]float64,w * h * 3),
    }
}

func (img *Image) SetPixel(x int, y int, color vec3.Vec3) {
    img.imageData[(y * img.width + x) * 3    ] = color.X;
    img.imageData[(y * img.width + x) * 3 + 1] = color.Y;
    img.imageData[(y * img.width + x) * 3 + 2] = color.Z;
}

func (img *Image) Write(filename string) error {
    writer := imageWriter.New( img.processImageData() , img.width, img.height)
    return writer.Write(filename);
}

func (i Image) processImageData() []float64 {
    out := make([]float64, i.width * i.height * 3)
    for i,v := range i.imageData {
        //Apply Gamma
        out[i] = math.Pow( v, 1 / gamma )
    }
    return out;
}