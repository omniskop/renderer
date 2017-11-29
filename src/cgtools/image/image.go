package image

import (
    "customtools/vec3"
    "cgtools/imageWriter"
)

const gamma = 2.2

type Image struct {
    width int
    height int
    imageData []float64
}

func New(w int, h int) Image {
    return Image{
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

func (img *Image) GetPixel(x,y int) vec3.Vec3 {
    return vec3.Vec3{
        img.imageData[(y * img.width + x) * 3    ],
        img.imageData[(y * img.width + x) * 3 + 1],
        img.imageData[(y * img.width + x) * 3 + 2],
    }
}

func (img *Image) Write(filename string) error {
    writer := imageWriter.New( img.imageData , img.width, img.height)
    return writer.Write(filename);
}

