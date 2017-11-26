package vec3

import (
    "math"
)

var White = Vec3{1,1,1}
var Black = Vec3{0,0,0}
var Green = Vec3{0,1,0}
var Blue = Vec3{0,0,1}
var Red = Vec3{1,0,0}
var Unit = Vec3{1,1,1}
var One = Vec3{1,1,1}
var Zero = Vec3{0,0,0} 
var Grey = Vec3{.5,.5,.5}
var Inf = Vec3{math.Inf(1), math.Inf(1), math.Inf(1)}

type Vec3 struct {
    X   float64
    Y   float64
    Z   float64
}

func New(x,y,z float64) Vec3 {
    return Vec3{x,y,z}
}

func (v Vec3) Length() float64 {
    return math.Sqrt( v.X * v.X + v.Y * v.Y + v.Z * v.Z )
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

func (v *Vec3) Normalize() {
    v.Divide(v.Length())
}

func (v Vec3) SquaredLength() float64 {
    return v.X * v.X + v.Y * v.Y + v.Z * v.Z
}

func (v Vec3) LessThan(a Vec3) bool {
    return v.X < a.X && v.Y < a.Y &&v.Z < a.Z
}

func (v Vec3) Equals(a Vec3) bool {
    return v.X == a.X && v.Y == a.Y && v.Z == a.Z
}

func (v *Vec3) Clamp() {
    v.X = math.Min(1, math.Max(v.X, 0))
    v.Y = math.Min(1, math.Max(v.Y, 0))
    v.Z = math.Min(1, math.Max(v.Z, 0))
}

func (v *Vec3) MultiplyByVec3(a Vec3) {
    v.X *= a.X
    v.Y *= a.Y
    v.Z *= a.Z
}



func Add(a Vec3, vecs ...Vec3) Vec3 {
    r := a.Clone()
    for _, vec := range vecs {
        r.Add(vec)
    }
    return r
}

func Subtract(a, b Vec3) Vec3 {
    return Vec3{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func MultiplyByVec3(a Vec3, vs ...Vec3) Vec3 {
    r := a.Clone()
    for _, vec := range vs {
        r.MultiplyByVec3(vec)
    }
    return r
}

func Multiply(n float64, v Vec3) Vec3 {
    return Vec3{ v.X * n, v.Y * n, v.Z * n }
}

func Divide(a Vec3, s float64) Vec3 {
    return Vec3{a.X / s, a.Y / s, a.Z / s}
}

func DivideVec3(a, b Vec3) Vec3 {
    return Vec3{
        a.X / b.X,
        a.Y / b.Y,
        a.Z / b.Z,
    }
}

func DotProduct(a,b Vec3) float64 {
    return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

func CrossProduct(a,b Vec3) Vec3 {
    return Vec3{a.Y * b.Z - a.Z * b.Y, a.Z * b.X - a.X * b.Z, a.X * b.Y - a.Y * b.X}
}

func Length(a Vec3) float64 {
    return a.Length()
}

func SquaredLength(a Vec3) float64 {
    return a.SquaredLength()
}

func Normalize(a Vec3) Vec3 {
    return Divide(a, a.Length())
}

func Project(a,b vec3.Vec3) vec3.Vec3 {
    c := vec3.Normalize(b)
    return vec3.Multiply( vec3.DotProduct(a, c),c)
}

func ProjectOnNormalized(a,b Vec3) Vec3 {
    return Multiply( DotProduct(a, b),b)
}

func Reject(a,b Vec3) Vec3 {
    return Subtract(
        a,
        Multiply(
            vec3.DotProduct(a,b) /
            vec3.DotProduct(b,b) ,
            b,
        ),
    )
}

func Clamp(v Vec3) Vec3 {
    return Vec3{math.Min(1, math.Max(v.X,0)), math.Min(1, math.Max(v.Y, 0)), math.Min(1, math.Max(v.Z, 0))}
}

func Hue(h float64) Vec3 {
    return Clamp(Vec3{
        math.Abs(h * 6 - 3) - 1,
        math.Abs(h * 6 - 2),
        math.Abs(h * 6 - 4),
    })
}

func HsvToRgb(hsv Vec3) Vec3 {
    return Multiply(hsv.Z, Add(Multiply(hsv.Y, Subtract(Hue(hsv.X), One)), One))
}

