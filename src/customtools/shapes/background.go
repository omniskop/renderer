package shapes

import (
    "math"
    "customtools/vec3"
    "customtools/ray"
)

type Background struct {
    Material    Material
}

func (this Background) Intersect(r ray.Ray) *Hit {
    return &Hit{
        T: math.Inf(1),
        Normal: vec3.Multiply(-1, r.Direction),
        Position: r.PointAt(math.Inf(1)),
        Material: this.Material,
    }
}