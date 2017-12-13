package shapes

import (
    "customtools/ray"
    "customtools/space"
)

type Group struct {
    Transformation      space.Transformation
    Shapes              []Shape
}

func (this Group) Intersect(r ray.Ray) *space.Hit {
    
    var closestHit *space.Hit
    var h *space.Hit
    r = this.Transformation.TransformRayIn(r)
    
    for _,shape := range this.Shapes {
        h = shape.Intersect(r)
        if h != nil && (closestHit == nil || h.T < closestHit.T) {
            closestHit = h
        }
    }
    
    if closestHit == nil {
        return nil
    }
    
    this.Transformation.TransformHitOut(closestHit)
    
    return closestHit
}