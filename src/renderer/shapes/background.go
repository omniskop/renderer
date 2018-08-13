package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// Background represents the background of the scene
type Background struct {
	Material space.Material
}

// Intersect returns the first hit of a ray with the object
func (bg Background) Intersect(r ray.Ray) *space.Hit {
	if !math.IsInf(r.T1, 0) {
		return nil
	}

	normal := vec3.Multiply(-1, r.Direction)

	inclination := -math.Acos(normal.Y)
	azimuth := math.Pi + math.Atan2(normal.X, normal.Z)
	u := (azimuth / (2 * math.Pi))
	v := inclination / math.Pi
	return &space.Hit{
		T:      math.Inf(1),
		Normal: normal,
		// Normal: vec3.Vec3{0,1,0},
		SurfaceCoordinates: vec3.Vec3{u, v, 0},
		Position:           r.PointAt(math.Inf(1)),
		Material:           bg.Material,
	}
}

// Includes checks if the point is inside the object
func (bg Background) Includes(point vec3.Vec3) bool {
	return false
}
