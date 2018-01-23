package texture

import (
	"customtools/vec3"
)

type textureColor struct {
	color vec3.Vec3
}

func NewColorVec(c vec3.Vec3) *textureColor {
	return &textureColor{c}
}

func NewColor(r, g, b float64) *textureColor {
	return &textureColor{vec3.Vec3{r, g, b}}
}

func (this *textureColor) SamplePoint(x, y, z float64) vec3.Vec3 {
	return this.color
}
