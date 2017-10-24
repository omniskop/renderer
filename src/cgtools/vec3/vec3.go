package vec3

import (
    
)

var White = Vec3{1,1,1}
var Green = Vec3{0,1,0}
var Blue = Vec3{0,0,1}
var Red = Vec3{1,0,0}
var Unit = Vec3{1,1,1}
var Zero = Vec3{0,0,0} 

type Vec3 struct {
    X   float64
    Y   float64
    Z   float64
}

func (v Vec3) Clone() Vec3 {
    return Vec3{v.X,v.Y,v.Z}
}

func (v *Vec3) Add(vec Vec3) {
    v.X += vec.X;
    v.Y += vec.Y;
    v.Z += vec.Z;
}

func (v *Vec3) Divide(n float64) {
    v.X /= n
    v.Y /= n
    v.Z /= n
}

func AddVec3(a Vec3, vecs ...Vec3) Vec3 {
    r := a.Clone()
    for _, vec := range vecs {
        r.Add(vec)
    }
    return r
}