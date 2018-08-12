package shapes

import (
	"log"
	"math"
	"space"
	"space/ray"
	"vec3"
)

// src https://www.uninformativ.de/bin/RaytracingSchnitttests-76a577a-CC-BY.pdf

type triangleMoller struct {
	v0       vec3.Vec3
	v1       vec3.Vec3
	v2       vec3.Vec3
	Material space.Material
}

func NewTriangleMoller(verticeA vec3.Vec3, verticeB vec3.Vec3, verticeC vec3.Vec3, m space.Material) triangleNew {
	log.Print(verticeA, verticeB, verticeC)
	return triangleNew{
		verticeA,
		verticeB,
		verticeC,
		m,
	}
}

func (this triangleMoller) Intersect(r ray.Ray) *space.Hit {
	v0v1 := vec3.Subtract(this.v1, this.v0)
	v0v2 := vec3.Subtract(this.v2, this.v0)
	pvec := vec3.CrossProduct(r.Direction, v0v2)
	det := vec3.DotProduct(v0v1, pvec)

	// ray and triangle are parallel if det is close to 0
	if math.Abs(det) < 0.00001 {
		return nil
	}
	invDet := 1.0000 / det

	tvec := vec3.Subtract(r.Origin, this.v0)
	u := vec3.DotProduct(tvec, pvec) * invDet
	if u < 0 || u > 1 {
		return nil
	}

	qvec := vec3.CrossProduct(tvec, v0v1)
	v := vec3.DotProduct(r.Direction, qvec) * invDet
	if v < 0 || u+v > 1 {
		return nil
	}

	t := vec3.DotProduct(v0v2, qvec) * invDet

	if t < r.T0 || t > r.T1 {
		return nil
	}

	return &space.Hit{
		T:                  t,
		Position:           r.PointAt(t),
		Normal:             vec3.Vec3{-1, 0, 0},
		SurfaceCoordinates: vec3.Vec3{0, 0, 0},
		Material:           this.Material,
	}
}

func (this triangleMoller) Includes(point vec3.Vec3) bool {
	return false
}
