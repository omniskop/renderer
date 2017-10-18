package vec3

import (
    
)

type Vec3 struct {
    X   float64
    Y   float64
    Z   float64
}

func (v Vec3) clone() Vec3 {
    return Vec3{v.X,v.Y,v.Z}
}

func (v Vec3) add(vec Vec3) {
    v.X += vec.X;
    v.Y += vec.Y;
    v.Z += vec.Z;
}

func addVec3(a Vec3, vecs ...Vec3) Vec3 {
    r := a.clone()
    for _, vec := range vecs {
        r.add(vec)
    }
    return r
}