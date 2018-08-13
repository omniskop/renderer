package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// Plane represents a plane
type Plane struct {
	Position vec3.Vec3
	Normal   vec3.Vec3
	Material space.Material
}

// Intersect returns the first hit of a ray with the object
func (plane Plane) Intersect(r ray.Ray) *space.Hit {
	t := vec3.DotProduct(vec3.Subtract(plane.Position, r.Origin), plane.Normal) / vec3.DotProduct(r.Direction, plane.Normal)

	if math.IsNaN(t) || t < r.T0 || t > r.T1 {
		return nil
	}

	var normal vec3.Vec3
	if vec3.DotProduct(plane.Normal, r.Direction) > 0 {
		normal = vec3.Multiply(-1, plane.Normal)
	} else {
		normal = plane.Normal
	}

	point := r.PointAt(t)

	return &space.Hit{
		T:                  t,
		Position:           point,
		Normal:             normal,
		SurfaceCoordinates: vec3.Vec3{point.Z, point.X, 0},
		Material:           plane.Material,
	}
}

// Includes checks if the point is inside the object
func (plane Plane) Includes(point vec3.Vec3) bool {
	return false
}
