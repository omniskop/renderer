package shapes

import (
	"space"
	"space/ray"
	"vec3"
)

// Group groups multiple shapes
type Group struct {
	Transformation space.Transformation
	Shapes         []Shape
}

// Intersect returns the first hit of a ray with the object
func (group Group) Intersect(r ray.Ray) *space.Hit {

	var closestHit *space.Hit
	var h *space.Hit
	r2 := group.Transformation.TransformRayIn(r)
	// r2 := r

	for _, shape := range group.Shapes {
		h = shape.Intersect(r2)
		if h != nil && (closestHit == nil || h.T < closestHit.T) {
			closestHit = h
		}
	}

	if closestHit == nil {
		return nil
	}

	group.Transformation.TransformHitOut(closestHit)

	return closestHit
}

// Includes checks if the point is inside the object
func (group Group) Includes(point vec3.Vec3) bool {
	for _, shape := range group.Shapes {
		if shape.Includes(point) {
			return true
		}
	}
	return false
}
