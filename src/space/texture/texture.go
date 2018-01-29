package texture

import (
	"vec3"
)

type Texture interface {
	SamplePoint(float64, float64, float64) vec3.Vec3
}
