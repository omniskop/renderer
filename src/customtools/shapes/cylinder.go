package shapes

import (
    "customtools/vec3"
    "customtools/ray"
    "math"
    "log"
)

type cylinder struct {
    Position    vec3.Vec3
    Normal   vec3.Vec3
    Radius      float64
    Height      float64
    Material    Material
    Disc1       Disc
    Disc2       Disc
    Plane       facingPlane
}

func NewCylinder(position, normal vec3.Vec3, radius, height float64, material Material) cylinder {
    out := cylinder{Position: position, Normal: normal, Radius: radius, Height: height, Material: material}
    out.Disc1 = Disc{position, normal, radius, material}
    out.Disc2 = Disc{vec3.Add(position, vec3.Multiply(height, normal)), normal, radius, material}
    out.Plane = facingPlane{position, normal, radius * 2, height}
    
    return out
}

func (this cylinder) Intersect(r ray.Ray) *Hit {
    
    disc1Hit := this.Disc1.Intersect(r)
    distanceToDisc1 := vec3.Subtract( r.Origin, this.Disc1.Position ).Length()
    disc2Hit := this.Disc2.Intersect(r)
    distanceToDisc2 := vec3.Subtract( r.Origin, this.Disc2.Position ).Length()
    
    
    if disc1Hit != nil && distanceToDisc1 < distanceToDisc2 {
        return disc1Hit
    }
    if disc2Hit != nil && distanceToDisc2 < distanceToDisc1 {
        return disc2Hit
    }
    
    var planeHit *Hit
    if disc1Hit == nil && disc2Hit == nil {
        planeHit = this.Plane.Intersect(r, true)
    } else {
        planeHit = this.Plane.Intersect(r, false)
    }
    
    if planeHit == nil {
        return nil
    }
    
    // if disc1Hit == nil && disc2Hit == nil {
    //     return &Hit{
    //         T: 1,
    //         Position: vec3.Vec3{0,0,0},
    //         Normal: vec3.Vec3{0,0,1},
    //         Material: Material_Diffuse{vec3.Green},
    //     }
    // } else if planeHit != nil {
    //     return &Hit{
    //         T: 1,
    //         Position: vec3.Vec3{0,0,0},
    //         Normal: vec3.Vec3{0,0,1},
    //         Material: Material_Diffuse{vec3.Red},
    //     }
    // }
    
    
    // The ray hit, now calculate correct hit Position
    //TODO: Multiply und Normaize müsste weggelassen werden können
    // centerToOuter := vec3.Multiply(this.Radius ,vec3.Normalize(vec3.CrossProduct(this.Normal, planeHit.Normal)))
    centerToOuter := vec3.CrossProduct(vec3.Subtract(this.Position, r.Origin), this.Normal )
    projectedHit := vec3.Project( vec3.Subtract(planeHit.Position ,this.Position), centerToOuter )
    reverseDirection := vec3.Multiply(-1, r.Direction)
    x := vec3.Project(reverseDirection, centerToOuter)// zur Breite der Ebene
    y := vec3.Project(reverseDirection, vec3.CrossProduct(centerToOuter, this.Normal))// zum Ray Ursprung in richtung Radius
    projectedDirection := vec3.Add(x,y)
    
    innerRay := ray.Ray{
        projectedHit,
        vec3.Normalize(projectedDirection),
        0,
        math.Inf(1),
    }
    // Für Die Kollision muss projectedDirection normalisiert werden, das T brauche ich aber relativ zum unnormalisierten
    
    sphere := Sphere{
        vec3.Vec3{0,0,0},
        this.Radius,
        Material_Diffuse{vec3.Red},
    }
    
    outerHit := sphere.Intersect(innerRay)
    
    if outerHit == nil {
        //Should not be possible but just in case
        return nil
    }
    
    if math.Abs(vec3.DotProduct(outerHit.Position, this.Normal)) > 0.0000001 {
        log.Print("error 2")
    }
    
    // geprüft: outerHit ist genau im Radius
    
    if math.Abs(outerHit.Position.Length() - this.Radius) > 0.0000001 {
        log.Print("error 1 ", outerHit.Position.Length())
    }
    
    // outerHit.Position isr korrekt
    
    // log.Print(outerHit.T)
    
    correctRay := ray.Ray{
        planeHit.Position,
        reverseDirection,
        0,
        math.Inf(1),
    }
    // Für Die Kollision muss projectedDirection normalisiert werden, das T brauche ich aber relativ zum unnormalisierten
    point := correctRay.PointAt(outerHit.T / (projectedDirection.Length() / vec3.Normalize(projectedDirection).Length()))
    // log.Print(vec3.Subtract(point ,this.Position), outerHit.Position)
    _ = log.Print
    
    x2 := vec3.Project(vec3.Subtract(point, this.Position), centerToOuter)// zur Breite der Ebene
    y2 := vec3.Project(vec3.Subtract(point, this.Position), vec3.CrossProduct(centerToOuter, this.Normal))// zum Ray Ursprung in richtung Radius
    projectedDirection2 := vec3.Add(x2,y2)
    
    if math.Abs(projectedDirection2.Length() - this.Radius) > 0.00000001 {
        log.Print(projectedDirection2.Length())
    }
    
    if math.Abs(vec3.DotProduct(outerHit.Normal, this.Normal)) > 0.0000001 {
        log.Print(outerHit.Normal, this.Normal)
    }
    
    // outerHit.Normal ist korrekt!
    return &Hit{
        T: vec3.Subtract(point, r.Origin).Length(),
        Position: point,
        Normal: outerHit.Normal,
        Material: this.Material,
        // Material: Material_Diffuse{
        //     outerHit.Normal,
        // },
    }
}

type facingPlane struct {
    Position    vec3.Vec3
    Upwards     vec3.Vec3
    Width       float64
    Height      float64
}

func (this facingPlane) Intersect(r ray.Ray, withHeightCheck bool) *Hit {
    normal := vec3.Normalize(vec3.CrossProduct(vec3.CrossProduct(vec3.Subtract(r.Origin, this.Position), this.Upwards), this.Upwards))
    t := vec3.DotProduct(vec3.Subtract(this.Position, r.Origin), normal) / vec3.DotProduct(r.Direction, normal)
    
    if math.IsNaN(t) || t < r.T0 || t > r.T1 {
        return nil
    }
    
    point := r.PointAt(t)
    
    positionToHit := vec3.Subtract(point, this.Position)
    
    // if  vec3.Reject(positionToHit, this.Upwards).Length() > this.Width / 2 || 
    //     vec3.DotProduct(positionToHit, this.Upwards) < 0 || vec3.Project(positionToHit, this.Upwards).Length() > this.Height {
    //     return nil
    // }
    
    if  vec3.Reject(positionToHit, this.Upwards).Length() > this.Width / 2 || 
        withHeightCheck && (vec3.DotProduct(positionToHit, this.Upwards) < 0 || vec3.Project(positionToHit, this.Upwards).Length() > this.Height) {
        return nil
    }
        
    return &Hit{
        T: t,
        Position: point,
        Normal: normal,
        Material: Material_Diffuse{vec3.Green},
    }
}











