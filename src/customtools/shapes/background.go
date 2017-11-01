package shapes

import (
    "math"
    "customtools/vec3"
    "customtools/ray"
    "customtools/hit"
)

type Background struct {
    Color       vec3.Vec3
}

func (this Background) Intersect(r ray.Ray) *hit.Hit {
    return &hit.Hit{
        T: math.Inf(1),
        Normal: vec3.Multiply(-1, r.Direction),
        Position: r.PointAt(math.Inf(1)),
        Color: this.Color,
    }
}