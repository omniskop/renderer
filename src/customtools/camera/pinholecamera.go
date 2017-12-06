package camera

import (
    "cgtools/mat4"
    "customtools/ray"
    "customtools/vec3"
    "math"
    "log"
)

type PinholeCamera struct {
    Position        vec3.Vec3
    Direction       vec3.Vec3
    Tilt            float64
    // transformationMatrix    mat4.mat4
    OpeningAngle    float64
    Width           int
    Height          int
}

// func NewPinholeCamera(position, direction, openingAngle, width, height) {
// 
//     return PinholeCamera{
//         position,
//         direction,
//         openingAngle,
//         width,
//         height
//     }
// }

func (this PinholeCamera) GetWidth() int {return this.Width}
func (this PinholeCamera) GetHeight() int {return this.Height}
func (this PinholeCamera) GetPosition() vec3.Vec3 {return this.Position}
func (this PinholeCamera) GetDirection() vec3.Vec3 {return this.Direction}
func (this PinholeCamera) GetOpeningAngle() float64 {return this.OpeningAngle}

func (this PinholeCamera) GetRayForPixel(x float64, y float64) ray.Ray {
    dir := vec3.Normalize(vec3.Vec3{
        X: x - float64(this.Width) / 2,
        Y: float64(this.Height / 2) - y,
        Z: -(float64(this.Width) / 2) / math.Tan( this.OpeningAngle / 2 ),
    })    
    
    transformationMatrix := mat4.NewTranslationByVec3(this.Position).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{0,0,1}, this.Tilt).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{0,1,0}, -math.Asin(this.Direction.X)).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{1,0,0}, math.Asin(this.Direction.Y)),
    )))
    // transformationMatrix := mat4.NewTranslationByVec3(this.Position)
    
    // horizontalAngle := -math.Asin(c.Direction.X)
    // heightAngle := math.Asin(c.Direction.Y)
    // horizontalRotationAxis := vec3.Vec3{0,1,0}
    // heightRotationAxis := vec3.CrossProduct(c.Direction, vec3.Vec3{0,1,0})
    // heightRotationAxis := vec3.Vec3{1,0,0}
    
    _ = log.Print
    //log.Print(transformationMatrix)
    
    cameraPosition := transformationMatrix.TransformPoint(vec3.Vec3{0,0,0})
    cameraDirection := transformationMatrix.TransformDirection(dir)
    
    out := ray.Ray{
        Origin: cameraPosition,
        Direction: cameraDirection,
        T0: 0,
        T1: math.Inf(1),
    }
    
    out.Direction.Normalize()
    
    return out
}