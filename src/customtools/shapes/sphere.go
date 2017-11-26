package shapes

import (
    "customtools/vec3"
    "customtools/ray"
    "math"
    "log"
)

type Sphere struct {
    Position    vec3.Vec3
    Radius      float64
    Material    Material
}

func (s Sphere) Intersect(r ray.Ray) *Hit {
    /*
        a = d^2
        b = 2x0d
        c = x0^2-r^2
    */
    newOrigin := vec3.Subtract( r.Origin, s.Position )
    
    a := vec3.DotProduct(r.Direction, r.Direction)
    b := 2 * vec3.DotProduct(newOrigin, r.Direction)
    c := vec3.DotProduct(newOrigin, newOrigin) - s.Radius * s.Radius
    
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
        if t1 < r.T0 || t1 > r.T1 {
            if t2 < r.T0 || t2 > r.T1 {
                return nil
            } else {
                offset = t2
            }
        } else {
            if t2 < r.T0 || t2 > r.T1 {
                offset = t1
            } else {
                offset = math.Min(t1, t2)
            }
        }
    }
        
    point := r.PointAt(offset)
    
    if math.Abs(vec3.Subtract(point, s.Position).Length() - s.Radius) > 0.000000001 {
        log.Print("fuck of ", offset, " ",vec3.Subtract(point, s.Position).Length() )
    }
    
    return &Hit{
        T: offset,
        Position: point,
        Normal: vec3.Divide( vec3.Subtract(point, s.Position), s.Radius ),
        Material: s.Material,
    }
}