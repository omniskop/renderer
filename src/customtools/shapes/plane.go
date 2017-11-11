package shapes

import (
    "math"
    "customtools/vec3"
    "customtools/ray"
)

type Plane struct {
    Position    vec3.Vec3
    Normal      vec3.Vec3
    Material    Material
}

func (this Plane) Intersect(r ray.Ray) *Hit {
    t := vec3.DotProduct(vec3.Subtract(this.Position, r.Origin), this.Normal) / vec3.DotProduct(r.Direction, this.Normal)
    
    if math.IsNaN(t) || t <= 0 {
        return nil
    }
    
    return &Hit{
        T: t,
        Position: r.PointAt(t),
        Normal: this.Normal,
        Material: this.Material,
    }
}