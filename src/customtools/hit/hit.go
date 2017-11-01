package hit

import (
    "customtools/vec3"
)

type Hit struct {
    T           float64
    Position    vec3.Vec3
    Normal      vec3.Vec3
    Color       vec3.Vec3
}