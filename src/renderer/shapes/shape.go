package shapes

import (
	"space"
	"space/ray"
	"vec3"
)

type Shape interface {
	Intersect(r ray.Ray) *space.Hit
	Includes(point vec3.Vec3) bool
}
