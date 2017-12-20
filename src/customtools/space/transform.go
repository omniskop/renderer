package space

import (
    "customtools/vec3"
    "customtools/ray"
    "cgtools/mat4"
    "log"
)

type Transformation struct {
    matrix          mat4.Mat4
    negatedMatrix   mat4.Mat4
    enabled         bool
}

func NoTransformation() Transformation {
    // return Transformation{mat4.NewIdentity()}
    return Transformation{enabled: false}
}

func NewTransformation(position vec3.Vec3, rotationX float64, rotationY float64, rotationZ float64) Transformation {
    k := mat4.NewTranslationByVec3(position).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{0,0,1}, rotationZ).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{0,1,0}, rotationY).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{1,0,0}, rotationX),
        )))
    k2 := mat4.NewTranslationByVec3(vec3.Multiply(-1,position)).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{0,0,1}, -rotationZ).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{0,1,0}, -rotationY).Multiply(
        mat4.NewRotationAroundVec3(vec3.Vec3{1,0,0}, -rotationX),
        )))
    return Transformation {
        matrix: k,
        negatedMatrix: k2,
        enabled: true,
    }
}

func negate(this mat4.Mat4) mat4.Mat4 {
    return mat4.Mat4{
        [16]float64{
            -this.Get(0,0),-this.Get(0,1),-this.Get(0,2),-this.Get(0,3),
            -this.Get(1,0),-this.Get(1,1),-this.Get(1,2),-this.Get(1,3),
            -this.Get(2,0),-this.Get(2,1),-this.Get(2,2),-this.Get(2,3),
            -this.Get(3,0),-this.Get(3,1),-this.Get(3,2),-this.Get(3,3),
        },
    }
}

func (this *Transformation) TransformRayIn(r ray.Ray) ray.Ray {
    if !this.enabled{return r}
    return ray.Ray{
        this.matrix.TransformPoint(r.Origin),
        // vec3.Add(this.matrix.GetPosition(), r.Origin),
        vec3.Normalize(this.matrix.TransformDirection(r.Direction)),
        r.T0,
        r.T1,
    }
}

// FIX FIRST
// func (this *Transformation) TransformPointIn(p vec3.Vec3) vec3.Vec3 {
//     if !this.enabled{return p}
//     return this.matrix.TransformPoint(p)
// }

// FIX FIRST
// func (this *Transformation) TransformHitIn(h *Hit) {
//     if !this.enabled{return}
//     h.Position = this.matrix.TransformPoint(h.Position)
//     h.Normal = this.matrix.TransformDirection(h.Normal)
// }


func (this *Transformation) TransformHitOut(h *Hit) {
    if !this.enabled{return}
    
    _ = log.Print
    if h.Normal.Y == 1 {
        log.Print("Rollmops")
    }
    
    // h.Position = this.matrix.TransformPoint(h.Position)
    h.Position = this.negatedMatrix.TransformPoint(h.Position)
    // h.Position = vec3.Add(this.negatedMatrix.GetPosition(), h.Position)
    h.Normal = this.negatedMatrix.TransformDirection(h.Normal)
}









