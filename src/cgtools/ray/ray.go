package ray

import (
    "cgtools/vec3"
)

type Ray struct {
    Origin  vec3.Vec3
    Direction vec3.Vec3
}

func (r Ray) PointAt(t float64) vec3.Vec3 {
    return vec3.Add(r.Origin, vec3.Multiply(t, r.Direction))
}