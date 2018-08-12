package shapes

import (
	"log"
	"math"
	"space"
	"space/ray"
	"vec3"
)

// src https://www.uninformativ.de/bin/RaytracingSchnitttests-76a577a-CC-BY.pdf

type triangleNew struct {
	v0       vec3.Vec3
	v1       vec3.Vec3
	v2       vec3.Vec3
	Material space.Material
}

func NewTriangleNew(verticeA vec3.Vec3, verticeB vec3.Vec3, verticeC vec3.Vec3, m space.Material) triangleNew {
	log.Print(verticeA, verticeB, verticeC)
	return triangleNew{
		verticeA,
		verticeB,
		verticeC,
		m,
	}
}

func (this triangleNew) Intersect(r ray.Ray) *space.Hit {
	// compute plane's normal
	v0v1 := vec3.Subtract(this.v1, this.v0)
	v0v2 := vec3.Subtract(this.v2, this.v0)
	// no need to normalize
	N := vec3.CrossProduct(v0v1, v0v2) // N
	// area2 := N.Length()

	// Step 1: finding P

	// check if ray and plane are parallel ?
	NdotRayDirection := vec3.DotProduct(N, r.Direction)
	if math.Abs(NdotRayDirection) < 0.000001 { // almost 0
		return nil // they are parallel so they don't intersect !
	}

	// compute d parameter using equation 2
	d := vec3.DotProduct(N, this.v0)

	// compute t (equation 3)
	t := (vec3.DotProduct(N, r.Origin) + d) / NdotRayDirection
	// check if the triangle is in behind the ray
	if t < r.T0 || t > r.T1 {
		return nil // the triangle is behind
	}

	// compute the intersection point using equation 1
	P := r.PointAt(t)

	// Step 2: inside-outside test
	var C vec3.Vec3 // vector perpendicular to triangle's plane

	// edge 0
	edge0 := vec3.Subtract(this.v1, this.v0)
	vp0 := vec3.Subtract(P, this.v0)
	C = vec3.CrossProduct(edge0, vp0)
	if vec3.DotProduct(N, C) < 0 {
		return nil // P is on the right side
	}

	// edge 1
	edge1 := vec3.Subtract(this.v2, this.v1)
	vp1 := vec3.Subtract(P, this.v1)
	C = vec3.CrossProduct(edge1, vp1)
	if vec3.DotProduct(N, C) < 0 {
		return nil // P is on the right side
	}

	// edge 2
	edge2 := vec3.Subtract(this.v0, this.v2)
	vp2 := vec3.Subtract(P, this.v2)
	C = vec3.CrossProduct(edge2, vp2)
	if vec3.DotProduct(N, C) < 0 {
		return nil // P is on the right side;
	}

	// this ray hits the triangle

	return &space.Hit{
		T:                  t,
		Position:           P,
		Normal:             N,
		SurfaceCoordinates: vec3.Vec3{0, 0, 0},
		Material:           this.Material,
	}
}

func (this triangleNew) Includes(point vec3.Vec3) bool {
	return false
}
