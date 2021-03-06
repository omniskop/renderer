package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// Sphere represents a sphere
type Sphere struct {
	Position vec3.Vec3
	Radius   float64
	Material space.Material
}

// Intersect returns the first hit of a ray with the object
func (s Sphere) Intersect(r ray.Ray) *space.Hit {
	/*
	   a = d^2
	   b = 2x0d
	   c = x0^2-r^2
	*/
	newOrigin := vec3.Subtract(r.Origin, s.Position)

	a := vec3.DotProduct(r.Direction, r.Direction)
	b := 2 * vec3.DotProduct(newOrigin, r.Direction)
	c := vec3.DotProduct(newOrigin, newOrigin) - s.Radius*s.Radius

	var offset float64

	switch n := b*b - 4*a*c; true {
	case n < 0:
		return nil
	case n == 0:
		offset =
			(-b +
				math.Sqrt(
					b*b-4*a*c,
				)) /
				2 * a
		if offset < 0 {
			return nil
		}
	case n > 0:
		t1 :=
			(-b +
				math.Sqrt(
					b*b-4*a*c,
				)) /
				2 * a
		t2 :=
			(-b -
				math.Sqrt(
					b*b-4*a*c,
				)) /
				2 * a
		if t1 < r.T0 || t1 > r.T1 {
			if t2 < r.T0 || t2 > r.T1 {
				return nil
			}
			offset = t2
		} else {
			if t2 < r.T0 || t2 > r.T1 {
				offset = t1
			} else {
				offset = math.Min(t1, t2)
			}
		}
	}

	point := r.PointAt(offset)

	normal := vec3.Divide(vec3.Subtract(point, s.Position), s.Radius)

	inclination := math.Acos(normal.Y)
	azimuth := math.Pi + math.Atan2(normal.X, normal.Z)
	u := azimuth / (2 * math.Pi)
	v := inclination / math.Pi

	return &space.Hit{
		T:                  offset,
		Position:           point,
		Normal:             normal,
		SurfaceCoordinates: vec3.Vec3{u, v, 0},
		Material:           s.Material,
	}
}

// Includes checks if the point is inside the object
func (s Sphere) Includes(point vec3.Vec3) bool {
	return vec3.Subtract(point, s.Position).SquaredLength() <= s.Radius
}
