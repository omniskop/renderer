package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// Disc represents a disc
type Disc struct {
	Position vec3.Vec3
	Normal   vec3.Vec3
	Radius   float64
	Material space.Material
}

// Intersect returns the first hit of a ray with the object
func (disc Disc) Intersect(r ray.Ray) *space.Hit {
	t := vec3.DotProduct(vec3.Subtract(disc.Position, r.Origin), disc.Normal) / vec3.DotProduct(r.Direction, disc.Normal)

	if math.IsNaN(t) || t < r.T0 || t > r.T1 || vec3.Subtract(r.PointAt(t), disc.Position).Length() > disc.Radius {
		return nil
	}

	var normal vec3.Vec3
	if vec3.DotProduct(disc.Normal, r.Direction) > 0 {
		normal = vec3.Multiply(-1, disc.Normal)
	} else {
		normal = disc.Normal
	}

	return &space.Hit{
		T:        t,
		Position: r.PointAt(t),
		Normal:   normal,
		Material: disc.Material,
	}
}

// Includes checks if the point is inside the object
func (disc Disc) Includes(point vec3.Vec3) bool {
	return false
}
