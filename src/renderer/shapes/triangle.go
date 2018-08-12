package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// src https://www.uninformativ.de/bin/RaytracingSchnitttests-76a577a-CC-BY.pdf

type triangle struct {
	n        vec3.Vec3
	uBeta    vec3.Vec3
	uGamma   vec3.Vec3
	d        float64
	kBeta    float64
	kGamma   float64
	Material space.Material
}

func NewTriangle(verticeA vec3.Vec3, verticeB vec3.Vec3, verticeC vec3.Vec3, m space.Material) triangle {
	out := triangle{}
	out.Material = m
	b := vec3.Subtract(verticeB, verticeA)
	c := vec3.Subtract(verticeC, verticeA)

	out.n = vec3.Normalize(vec3.CrossProduct(b, c))
	out.d = vec3.DotProduct(out.n, verticeA)

	bb := vec3.DotProduct(b, b)
	bc := vec3.DotProduct(b, c)
	cc := vec3.DotProduct(c, c)

	D := 1.0 / (cc*bb - bc*bc)
	bbD := bb * D
	bcD := bc * D
	ccD := cc * D

	out.uBeta = vec3.Subtract(vec3.Multiply(ccD, b), vec3.Multiply(bcD, c))
	out.uGamma = vec3.Subtract(vec3.Multiply(bbD, c), vec3.Multiply(bcD, b))

	out.kBeta = -vec3.DotProduct(verticeA, out.uBeta)
	out.kGamma = -vec3.DotProduct(verticeA, out.uGamma)

	return out
}

func (this triangle) Intersect(r ray.Ray) *space.Hit {
	rn := vec3.DotProduct(r.Direction, this.n)
	if math.Abs(rn) < 0.000001 {
		return nil
	}

	alpha1 := (d - vec3.DotProduct(r.Origin, this.n)) / rn
	if alpha1 < r.T0 || alpha1 > r.T1 {
		return nil
	}

	q := r.PointAt(alpha1)
	beta := vec3.DotProduct(this.uBeta, q) + this.kBeta
	if beta < 0 {
		return nil
	}

	gamma := vec3.DotProduct(this.uGamma, q) + this.kGamma
	if gamma < 0 {
		return nil
	}

	alpha := 1 - beta - gamma
	if alpha < 0 {
		return nil
	}

	return &space.Hit{
		T:                  alpha1,
		Position:           q,
		Normal:             this.n,
		SurfaceCoordinates: vec3.Vec3{0, 0, 0},
		Material:           this.Material,
	}
}

func (this triangle) Includes(point vec3.Vec3) bool {
	return false
}
