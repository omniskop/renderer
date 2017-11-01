package shapes

import (
    "customtools/vec3"
    "customtools/ray"
    "customtools/hit"
    "math"
)

type Sphere struct {
    Position    vec3.Vec3
    Radius      int
}

func (s Sphere) Intersect(r ray.Ray) *hit.Hit {
    /*
        a = d^2
        b = 2x0d
        c = x0^2-r^2
    */
    newOrigin := vec3.Subtract( r.Origin, s.Position )
    
    a := vec3.DotProduct(r.Direction, r.Direction)
    b := 2 * vec3.DotProduct(newOrigin, r.Direction)
    c := vec3.DotProduct(newOrigin, newOrigin) - float64(s.Radius * s.Radius)
    
    var offset float64
    
    switch n := b * b - 4 * a * c; true {
    case n < 0:
        return nil
    case n == 0:
        offset = 
            (-b +
            math.Sqrt(
                b * b - 4 * a * c,
            )) /
            2 * a
        if offset < 0 {
            return nil
        }
    case n > 0:
        t1 :=
            (-b +
            math.Sqrt(
                b * b - 4 * a * c,
            )) /
            2 * a
        t2 :=
            (-b -
            math.Sqrt(
                b * b - 4 * a * c,
            )) /
            2 * a
        if t1 < 0 {
            if t2 < 0 {
                return nil
            } else {
                offset = t2
            }
        } else {
            if t2 < 0 {
                offset = t1
            } else {
                offset = math.Min(t1, t2)
            }
        }
    }
        
    point := r.PointAt(offset)
    
    return &hit.Hit{
        T: r,
        Position: point,
        Normal: vec3.Divide( vec3.Subtract(point, s.Position), float64(s.Radius) ),
    }
}