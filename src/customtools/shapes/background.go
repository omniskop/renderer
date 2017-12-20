package shapes

import (
    "math"
    "customtools/vec3"
    "customtools/ray"
    "customtools/space"
)

type Background struct {
    Material    space.Material
}

func (this Background) Intersect(r ray.Ray) *space.Hit {
    return &space.Hit{
        T: math.Inf(1),
        // Normal: vec3.Multiply(-1, r.Direction),
        Normal: vec3.Vec3{0,1,0},
        Position: r.PointAt(math.Inf(1)),
        Material: this.Material,
    }
}