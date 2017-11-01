package camera

import (
    "cgtools/ray"
    "cgtools/vec3"
    "math"
)

type Camera struct {
    OpeningAngle    float64
    Width           int
    Height          int
}

// func (c Camera) GetRayForPixel(x int, y int) ray.Ray {
//     xProjectionNormalPlane := x / c.Width * math.Tan( OpeningAngle / 2 )*2
//     yProjectionNormalPlane := y / c.Height * math.Tan( OpeningAngle / 2 )*2
//     
//     xAngle = math.Atan( xProjectionNormalPlane )
//     yAngle = math.Atan( yProjectionNormalPlane )
//     
//     out := vec3.Vec3{0,0,0}
//     
//     out.X = 
// }

func (c Camera) GetRayForPixel(x int, y int) ray.Ray {
    out := ray.Ray{Origin: vec3.Zero, Direction: vec3.Vec3{
        X: float64(x) - float64(c.Width) / 2,
        Y: float64(c.Height / 2) - float64(y),
        Z: -(float64(c.Width) / 2) / math.Tan( c.OpeningAngle / 2 ),
    }}
    
    out.Direction.Normalize()
    
    return out
}