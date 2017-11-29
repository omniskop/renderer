package camera

import (
    "customtools/ray"
    "customtools/vec3"
    "math"
    "log"
)

type PinholeCamera struct {
    Position        vec3.Vec3
    Direction       vec3.Vec3
    OpeningAngle    float64
    Width           int
    Height          int
}

func (this PinholeCamera) GetWidth() int {return this.Width}
func (this PinholeCamera) GetHeight() int {return this.Height}
func (this PinholeCamera) GetPosition() vec3.Vec3 {return this.Position}
func (this PinholeCamera) GetDirection() vec3.Vec3 {return this.Direction}
func (this PinholeCamera) GetOpeningAngle() float64 {return this.OpeningAngle}

func (c PinholeCamera) GetRayForPixel(x float64, y float64) ray.Ray {
    dir := vec3.Normalize(vec3.Vec3{
        X: x - float64(c.Width) / 2,
        Y: float64(c.Height / 2) - y,
        Z: -(float64(c.Width) / 2) / math.Tan( c.OpeningAngle / 2 ),
    })
    
    // dir = vec3.Vec3{0,-math.Cos(math.Pi / 4), -math.Cos(math.Pi / 4)}
    
    horizontalAngle := -math.Asin(c.Direction.X)
    heightAngle := math.Asin(c.Direction.Y)
    horizontalRotationAxis := vec3.Vec3{0,1,0}
    heightRotationAxis := vec3.CrossProduct(c.Direction, vec3.Vec3{0,1,0})
    // heightRotationAxis := vec3.Vec3{1,0,0}
    
    _ = log.Print
    // log.Print("begin")
    // log.Print(horizontalAngle)
    // log.Print(heightAngle)
    // log.Print(horizontalRotationAxis)
    // log.Print(heightRotationAxis)
    
    k := vec3.Vec3{0,0,0}
        
    k = vec3.Add(
        vec3.Multiply(math.Cos(horizontalAngle), dir),
        vec3.Multiply(math.Sin(horizontalAngle), vec3.CrossProduct(horizontalRotationAxis, dir)),
        vec3.Multiply((1-math.Cos(horizontalAngle)) * vec3.DotProduct(horizontalRotationAxis, dir),horizontalRotationAxis),
    )
    
    // log.Print(k)
    
    // log.Print(math.Cos(horizontalAngle))
    
    // a := (vec3.Multiply(math.Cos(heightAngle), k)).Z
    // b := (vec3.Multiply(math.Sin(heightAngle), vec3.CrossProduct(heightRotationAxis, k))).Z
    // d := (vec3.Multiply((1-math.Cos(heightAngle)) * vec3.DotProduct(heightRotationAxis, k),heightRotationAxis)).Z

    // log.Print(a,b,d)
    // log.Print(a+b+d)
    
    k = vec3.Add(
        vec3.Multiply(math.Cos(heightAngle), k),
        vec3.Multiply(math.Sin(heightAngle), vec3.CrossProduct(heightRotationAxis, k)),
        vec3.Multiply((1-math.Cos(heightAngle)) * vec3.DotProduct(heightRotationAxis, k),heightRotationAxis),
    )
    
    // log.Print(k)
    
    out := ray.Ray{
        Origin: c.Position,
        Direction: k,
        T0: 0,
        T1: math.Inf(1),
    }
    
    out.Direction.Normalize()
    
    return out
}