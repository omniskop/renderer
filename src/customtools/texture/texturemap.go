package texture

import (
	"cgtools/imageTexture"
	"customtools/vec3"
)

type texturemap struct {
	image   *imageTexture.ImageTexture
	scaling float64
}

func NewTexturemap(filename string, gamma float64) *texturemap {
	out := texturemap{}
	out.image = imageTexture.New(filename, gamma)
	out.scaling = 1
	return &out
}

func NewScaledTexturemap(filename string, gamma float64, scaling float64) *texturemap {
	out := texturemap{}
	out.image = imageTexture.New(filename, gamma)
	out.scaling = scaling
	return &out
}

func (this *texturemap) SamplePoint(x, y, z float64) vec3.Vec3 {
	return this.image.SamplePoint(x*this.scaling, y*this.scaling)
}
