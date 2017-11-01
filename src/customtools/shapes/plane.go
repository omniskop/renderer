package shapes

import (
    "math"
    "customtools/vec3"
    "customtools/ray"
    "customtools/hit"
)

type Plane struct {
    Position    vec3.Vec3
    Normal      vec3.Vec3
    Color       vec3.Vec3
}

func (this Plane) Intersect(r ray.Ray) *hit.Hit {
    t := vec3.DotProduct(vec3.Subtract(this.Position, r.Origin), this.Normal) / vec3.DotProduct(r.Direction, this.Normal)
    
    if math.IsNaN(t) || t < 0 {
        return nil
    }
    
    return &hit.Hit{
        T: t,
        Position: r.PointAt(t),
        Normal: this.Normal,
        Color: this.Color,
    }
}