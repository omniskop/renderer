package shapes

import (
	"customtools/ray"
	"customtools/space"
	"customtools/texture"
	"customtools/vec3"
)

// ======  UNION  ======

func Union(elements ...Shape) Shape {
	return Group{space.NoTransformation(), elements}
}

// ======  DIFFERENCE  ======

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

func (this differenceGroup) Intersect(r ray.Ray) *space.Hit {
	hit := this.minuend.Intersect(r)
	if hit == nil {
		return nil
	}
	hit.Material = space.Material_Sky{texture.NewColor(1, 0, 0)}
	if this.subtrahend.Includes(hit.Position) {
		orignalT := hit.T
		r.Origin = hit.Position
		// r.T0 = 0.1
		hit = this.subtrahend.Intersect(r)
		if hit == nil { // Should not happen. But it can when t.T0 is not 0
			return nil
		}
		if !this.minuend.Includes(hit.Position) {
			return nil
		}
		hit.Normal = vec3.Multiply(-1, hit.Normal)
		hit.T += orignalT
	}
	return hit
}

func (this differenceGroup) Includes(point vec3.Vec3) bool {
	return this.minuend.Includes(point) && !this.subtrahend.Includes(point)
}

// ======  INTERSECTION  ======

func Intersection(shapes ...Shape) Shape {
	return intersectionGroup{
		shapes: shapes,
	}
}

type intersectionGroup struct {
	shapes []Shape
}

func (this intersectionGroup) Intersect(r ray.Ray) *space.Hit {
	var furthestHit *space.Hit = nil
	var h *space.Hit

	for _, shape := range this.shapes {
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

func (this intersectionGroup) Includes(point vec3.Vec3) bool {
	for _, shape := range this.shapes {
		if !shape.Includes(point) {
			return false
		}
	}
	return true
}
