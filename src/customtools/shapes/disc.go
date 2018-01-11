package shapes

import (
    "customtools/vec3"
    "customtools/ray"
    "customtools/space"
    "math"
)

type Disc struct {
    Position    vec3.Vec3
    Normal   vec3.Vec3
    Radius      float64
    Material    space.Material
}

func (this Disc) Intersect(r ray.Ray) *space.Hit {
    t := vec3.DotProduct(vec3.Subtract(this.Position, r.Origin), this.Normal) / vec3.DotProduct(r.Direction, this.Normal)
    
    if math.IsNaN(t) || t < r.T0 || t > r.T1 || vec3.Subtract(r.PointAt(t), this.Position).Length() > this.Radius {
        return nil
    }
    
    var normal vec3.Vec3
    if vec3.DotProduct(this.Normal, r.Direction) > 0 {
        normal = vec3.Multiply(-1,this.Normal)
    } else {
        normal = this.Normal
    }
    
    return &space.Hit{
        T: t,
        Position: r.PointAt(t),
        Normal: normal,
        Material: this.Material,
    }
}

func (this Disc) Includes(point vec3.Vec3) bool {
	return false
}