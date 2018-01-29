package ray

import (
	"vec3"
)

type Ray struct {
	Origin    vec3.Vec3
	Direction vec3.Vec3
	T0        float64
	T1        float64
}

func (r Ray) PointAt(t float64) vec3.Vec3 {
	return vec3.Add(r.Origin, vec3.Multiply(t, r.Direction))
}

func (r Ray) InRange(t float64) bool {
	return t > r.T0 && t < r.T1
}
