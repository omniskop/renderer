package shapes

import (
    "customtools/ray"
    "customtools/space"
    _"customtools/vec3"
    _"log"
    _"runtime"
    _"math"
)

type Group struct {
    Transformation      space.Transformation
    Shapes              []Shape
}

func (this Group) Intersect(r ray.Ray) *space.Hit {
    
    var closestHit *space.Hit
    var h *space.Hit
    // r2 := this.Transformation.TransformRayIn(r)
    r2 := r
    
    for _,shape := range this.Shapes {
        h = shape.Intersect(r2)
        if h != nil && (closestHit == nil || h.T < closestHit.T) {
            closestHit = h
        }
    }
    
    if closestHit == nil {
        return nil
    }
    
    // prev := vec3.Vec3{closestHit.Normal.X, closestHit.Normal.Y, closestHit.Normal.Z}
    
    // this.Transformation.TransformHitOut(closestHit)
    
    // if len(this.Shapes) == 2 && closestHit.Position.X == math.Inf(1) {
    //     runtime.Breakpoint()
    // }
    // 
    // if false && len(this.Shapes) == 2 && closestHit.Normal.Equals(vec3.Vec3{-0,-1,-0}) {
    //     log.Print("There they are ", prev, closestHit.Material)
    // }
    // 
    // _, e := closestHit.Material.(space.Material_Diffuse)
    // if false && len(this.Shapes) == 1 && e && r.PointAt(closestHit.T).Y < -0.01 {
    //     log.Print("Der Übeltäter ")
    // }
    
    return closestHit
}