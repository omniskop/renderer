package shapes

import (
	"space"
	"space/ray"
	"vec3"
)

// A Shape offers a common interface for geometry
type Shape interface {
	Intersect(r ray.Ray) *space.Hit
	Includes(point vec3.Vec3) bool
}
