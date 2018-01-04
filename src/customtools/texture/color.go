package texture;

import (
    "customtools/vec3"
)

type textureColor struct {
    color       vec3.Vec3
}

func NewColor(c vec3.Vec3) *textureColor {
    return &textureColor{c}
}

func (this *textureColor) SamplePoint(x,y float64) vec3.Vec3 {
    return this.color
}