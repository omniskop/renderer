package shapes

import (
    "customtools/ray"
)

type Group struct {
    Shapes      []Shape
}

func (this Group) Intersect(r ray.Ray) *Hit {
    
    var closestHit *Hit
    var h *Hit
    
    for _,shape := range this.Shapes {
        h = shape.Intersect(r)
        if h != nil && (closestHit == nil || h.T < closestHit.T) {
            closestHit = h
        }
    }
    
    if closestHit == nil {
        return nil
    }
    
    return closestHit
}