package shapes

import (
	"customtools/ray"
	"customtools/space"
	"customtools/vec3"
	"math"
)

type Background struct {
	Material space.Material
}

func (this Background) Intersect(r ray.Ray) *space.Hit {
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
		Material:           this.Material,
	}
}

func (this Background) Includes(point vec3.Vec3) bool {
	return false
}
