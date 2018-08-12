// +build ignore

package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// src https://www.uninformativ.de/bin/RaytracingSchnitttests-76a577a-CC-BY.pdf

type triangleHavel struct {
	n0
	n1
	n2 vec3.Vec3
	d0
	d1
	d2       float64
	Material space.Material
}

func NewTriangleHavel(verticeA vec3.Vec3, verticeB vec3.Vec3, verticeC vec3.Vec3, m space.Material) triangleNew {
	out := triangleHavel{Material: m}
	e1 := vec3.Subtract(verticeB, verticeA)
	e2 := vec3.Subtract(verticeC, verticeA)

	// D->n0 = v3_cross(e1, e2);
	out.n0 = vec3.CrossProduct(e1, e2)
	// D->d0 = v3_dot(D->n0, v0);
	out.d0 = vec3.DotProduct(out.n0, verticeA)

	// float inv_denom = 1 / v3_dot(D->n0, D->n0);
	inv_denom := 1.0 / vec3.DotProduct(out.n0, out.n0)

	// D->n1 = v3_scale(v3_cross(e2, D->n0), inv_denom);
	out.n1 = vec3.Multiply(inv_denom, vec3.CrossProduct(e2, out.n0))
	// D->d1 = -v3_dot(D->n1, v0);
	out.d1 = -vec3.DotProduct(D.n1, verticeA)

	// D->n2 = v3_scale(v3_cross(D->n0, e1), inv_denom);
	out.n2 = vec3.Multiply(inv_denom, vec3.CrossProduct(out.n0, e1))
	// D->d2 = -v3_dot(D->n2, v0);
	out.d2 = -vec3.DotProduct(out.n2, verticeA)

	return out
}

func (this triangleHavel) Intersect(r ray.Ray) *space.Hit {
	// float det = v3_dot(D->n0, d);
	det := vec3.DotProduct(this.n0, r.Direction)
	// float dett = D->d0 - v3_dot(o, D->n0);
	dett := this.d0 - vec3.DotProduct(o, this.n0)
	// vec3 wr = v3_add(v3_scale(o, det), v3_scale(d, dett));
	wr := vec3.Add(vec3.Multiply(det, r.Origin), vec3.Multiply(dett, r,Direction))
	// uv->x = v3_dot(wr, D->n1) + det * D->d1;
	// uv->y = v3_dot(wr, D->n2) + det * D->d2;
	uv := vec3.Vec3{
		vec3.DotProduct(wr, this.n1) + det * this.d1,
		vec3.DotProduct(wr, this.n2) + det * this.d2,
		0,
	}
	// float tmpdet0 = det - uv->x - uv->y;
	tmpdet0 := det - uv.X - uv.Y
	// int pdet0 = ((int_or_float)tmpdet0).i;
	pdet0 := int(tmpdet0)
	// int pdetu = ((int_or_float)uv->x).i;
	pdetu := int(uv.X)
	// int pdetv = ((int_or_float)uv->y).i;
	pdetv := int(uv.Y)
	// pdet0 = pdet0 ^ pdetu;
	
    pdet0 = pdet0 | (pdetu ^ pdetv);
    if (pdet0 & 0x80000000)
        return false;
    float rdet = 1 / det;
    uv->x *= rdet;
    uv->y *= rdet;
    *t = dett * rdet;
    return *t >= ISECT_NEAR && *t <= ISECT_FAR;


	return &space.Hit{
		T:                  t,
		Position:           r.PointAt(t),
		Normal:             vec3.Vec3{-1, 0, 0},
		SurfaceCoordinates: vec3.Vec3{0, 0, 0},
		Material:           this.Material,
	}
}

func (this triangleHavel) Includes(point vec3.Vec3) bool {
	return false
}
