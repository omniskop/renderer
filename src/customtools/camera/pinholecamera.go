package camera

import (
    "customtools/ray"
    "customtools/vec3"
    "math"
)

type PinholeCamera struct {
    Position        vec3.Vec3
    Direction       vec3.Vec3
    OpeningAngle    float64
    Width           int
    Height          int
}

func (c PinholeCamera) GetRayForPixel(x float64, y float64) ray.Ray {
    dir := vec3.Normalize(vec3.Vec3{
        X: x - float64(c.Width) / 2,
        Y: float64(c.Height / 2) - y,
        Z: -(float64(c.Width) / 2) / math.Tan( c.OpeningAngle / 2 ),
    })
    
    horizontalAngle := -math.Asin(c.Direction.X)
    heightAngle := math.Asin(c.Direction.Y)
    horizontalRotationAxis := vec3.Vec3{0,1,0}
    heightRotationAxis := vec3.CrossProduct(c.Direction, vec3.Vec3{0,1,0})
    
    k := vec3.Vec3{0,0,0}
        
    k = vec3.Add(
        vec3.Multiply(math.Cos(horizontalAngle), dir),
        vec3.Multiply(math.Sin(horizontalAngle), vec3.CrossProduct(horizontalRotationAxis, dir)),
        vec3.Multiply((1-math.Cos(horizontalAngle)) * vec3.DotProduct(horizontalRotationAxis, dir),horizontalRotationAxis),
    )
    
    k = vec3.Add(
        vec3.Multiply(math.Cos(heightAngle), k),
        vec3.Multiply(math.Sin(heightAngle), vec3.CrossProduct(heightRotationAxis, k)),
        vec3.Multiply((1-math.Cos(heightAngle)) * vec3.DotProduct(heightRotationAxis, k),heightRotationAxis),
    )
    
    out := ray.Ray{
        Origin: c.Position,
        Direction: k,
        T0: 0,
        T1: math.Inf(1),
    }
    
    out.Direction.Normalize()
    
    return out
}