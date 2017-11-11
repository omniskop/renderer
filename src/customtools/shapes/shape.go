package shapes

import (
    "customtools/ray"
    "customtools/vec3"
)

type Shape interface {
    Intersect(r ray.Ray) *Hit
}

type Hit struct {
    T           float64
    Position    vec3.Vec3
    Normal      vec3.Vec3
    Material    Material
}