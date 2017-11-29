package camera

import (
    "customtools/ray"
    "customtools/vec3"
)

type Camera interface {
    GetRayForPixel(x float64, y float64) ray.Ray
    GetWidth()   int
    GetHeight()   int
    GetPosition()   vec3.Vec3
    GetDirection()   vec3.Vec3
    GetOpeningAngle()   float64
}