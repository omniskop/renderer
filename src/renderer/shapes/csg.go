package shapes

import (
	"space"
	"space/ray"
	"vec3"
)

// ======  UNION  ======

// Union combines the geometry of multiple shapes
func Union(elements ...Shape) Shape {
	return Group{space.NoTransformation(), elements}
}

// ======  DIFFERENCE  ======

// Difference substracts the geometry multiple shapes from a single shape
func Difference(mainShape Shape, shapes ...Shape) Shape {
	return differenceGroup{
		minuend:    mainShape,
		subtrahend: Group{space.NoTransformation(), shapes},
	}
}

type differenceGroup struct {
	minuend    Shape // main shape
	subtrahend Shape
}

// Intersect returns the first hit of a ray with the object
func (diffGroup differenceGroup) Intersect(r ray.Ray) *space.Hit {
	hit := diffGroup.minuend.Intersect(r)
	if hit == nil {
		return nil
	}
	if diffGroup.subtrahend.Includes(hit.Position) {
		orignalT := hit.T
		r.Origin = hit.Position
		// r.T0 = 0.1
		hit = diffGroup.subtrahend.Intersect(r)
		if hit == nil { // Should not happen. But it can when t.T0 is not 0
			return nil
		}
		if !diffGroup.minuend.Includes(hit.Position) {
			return nil
		}
		hit.Normal = vec3.Multiply(-1, hit.Normal)
		hit.T += orignalT
	}
	return hit
}

// Includes checks if the point is inside the object
func (diffGroup differenceGroup) Includes(point vec3.Vec3) bool {
	return diffGroup.minuend.Includes(point) && !diffGroup.subtrahend.Includes(point)
}

// ======  INTERSECTION  ======

// Intersection represents an and on all shape geometry
func Intersection(shapes ...Shape) Shape {
	return intersectionGroup{
		shapes: shapes,
	}
}

type intersectionGroup struct {
	shapes []Shape
}

// Intersect returns the first hit of a ray with the object
func (intGroup intersectionGroup) Intersect(r ray.Ray) *space.Hit {
	var furthestHit *space.Hit
	var h *space.Hit

	for _, shape := range intGroup.shapes {
		h = shape.Intersect(r)
		if h == nil {
			return nil
		}
		if furthestHit == nil || h.T > furthestHit.T {
			furthestHit = h
		}
	}

	return furthestHit
}

// Includes checks if the point is inside the object
func (intGroup intersectionGroup) Includes(point vec3.Vec3) bool {
	for _, shape := range intGroup.shapes {
		if !shape.Includes(point) {
			return false
		}
	}
	return true
}
