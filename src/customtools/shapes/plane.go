package shapes

import (
	"customtools/ray"
	"customtools/space"
	"customtools/vec3"
	"log"
	"math"
)

type Plane struct {
	Position vec3.Vec3
	Normal   vec3.Vec3
	Material space.Material
}

func (this Plane) Intersect(r ray.Ray) *space.Hit {
	t := vec3.DotProduct(vec3.Subtract(this.Position, r.Origin), this.Normal) / vec3.DotProduct(r.Direction, this.Normal)

	if math.IsNaN(t) || t < r.T0 || t > r.T1 {
		return nil
	}

	var normal vec3.Vec3
	if vec3.DotProduct(this.Normal, r.Direction) > 0 {
		// log.Print("nope ", r.Direction.Length(), r.Origin)
		_ = log.Print
		normal = vec3.Multiply(-1, this.Normal)
	} else {
		normal = this.Normal
	}

	point := r.PointAt(t)

	return &space.Hit{
		T:                  t,
		Position:           point,
		Normal:             normal,
		SurfaceCoordinates: vec3.Vec3{point.Z, point.X, 0},
		Material:           this.Material,
	}
}
