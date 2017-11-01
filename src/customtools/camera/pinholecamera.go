package camera

import (
    "customtools/ray"
    "customtools/vec3"
    "math"
)

type PinholeCamera struct {
    OpeningAngle    float64
    Width           int
    Height          int
}

func (c PinholeCamera) GetRayForPixel(x int, y int) ray.Ray {
    out := ray.Ray{Origin: vec3.Zero, Direction: vec3.Vec3{
        X: float64(x) - float64(c.Width) / 2,
        Y: float64(c.Height / 2) - float64(y),
        Z: -(float64(c.Width) / 2) / math.Tan( c.OpeningAngle / 2 ),
    }}
    
    out.Direction.Normalize()
    
    return out
}