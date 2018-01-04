package space

import (
    "customtools/vec3"
)

type Hit struct {
    T           float64
    Position    vec3.Vec3
    Normal      vec3.Vec3
    SurfaceCoordinates  vec3.Vec3
    Material    Material
}