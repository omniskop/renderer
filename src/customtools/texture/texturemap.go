package texture;

import (
    "cgtools/imageTexture"
    "customtools/vec3"
)

type texturemap struct {
    image       *imageTexture.ImageTexture
}

func NewTexturemap(filename string, gamma float64) *texturemap {
    out := texturemap{}
    out.image = imageTexture.New(filename, gamma)
    return &out
}

func (this *texturemap) SamplePoint(x,y float64) vec3.Vec3 {
    return this.image.SamplePoint(x,y)
}