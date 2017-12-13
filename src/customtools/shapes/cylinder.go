package shapes

import (
    "customtools/vec3"
    "customtools/ray"
    "customtools/space"
)

type Cylinder struct {
    Position    vec3.Vec3
    Normal   vec3.Vec3
    Radius      float64
    Height      float64
    Material    space.Material
    disc1       Disc
    disc2       Disc
    walls       HollowCylinder
}

func NewCylinder(position, normal vec3.Vec3, radius, height float64, material space.Material) Cylinder {
    out := Cylinder{Position: position, Normal: normal, Radius: radius, Height: height, Material: material}
    out.disc1 = Disc{position, vec3.Multiply(-1,normal), radius, material}
    out.disc2 = Disc{vec3.Add(position, vec3.Multiply(height, normal)), normal, radius, material}
    out.walls = NewHollowCylinder(position, normal, radius, height, material)
    
    return out
}

func (this Cylinder) Intersect(r ray.Ray) *space.Hit {
    if vec3.DotProduct(this.disc1.Normal, r.Direction) < 0 {
        disc1Hit := this.disc1.Intersect(r)
        
        if disc1Hit != nil {
            return disc1Hit
        }
    }
    if vec3.DotProduct(this.disc2.Normal, r.Direction) < 0 {
        disc2Hit := this.disc2.Intersect(r)
        if disc2Hit != nil {
            return disc2Hit
        }
    }
    
    cylinderHit := this.walls.Intersect(r)
    
    return cylinderHit
}

type HollowCylinder struct {
    Position    vec3.Vec3
    Normal   vec3.Vec3
    Radius      float64
    Height      float64
    Material    space.Material
    collisionSphere Sphere
}

func NewHollowCylinder(position, normal vec3.Vec3, radius, height float64, material space.Material) HollowCylinder {
    out := HollowCylinder{Position: position, Normal: normal, Radius: radius, Height: height, Material: material}
    out.collisionSphere = Sphere{
        Position: vec3.Vec3{0,0,0},
        Radius: radius,
    }
    return out
}

func (this HollowCylinder) Intersect(r ray.Ray) *space.Hit {
    xAxis := vec3.Normalize(vec3.CrossProduct(this.Normal, r.Direction)) // Durch die Direction kann man sich eine Projektion für die Richtung spaaren
    if xAxis.Equals(vec3.Zero) {
        // Happens when this.Normal and r.Direction are equal.
        // This is very unlikely but it can happen ans will completly blackout the image
        return nil
    }
    yAxis := vec3.Normalize(vec3.CrossProduct(this.Normal, xAxis))
    positionToOrigin := vec3.Subtract(r.Origin,this.Position)
    
    projectedOrigin := vec3.Add(
        vec3.Project(positionToOrigin, xAxis),
        vec3.Project(positionToOrigin, yAxis),
    )
    // The xAxis can be left away due to it being calculated bases on the
    // direction of the ray so the projection would just result in a zero-vector
    projectedDirection :=
        vec3.Project(r.Direction, yAxis)
    
    // The direction of the ray has to be normalized because of that I need do
    // scale the T of the hit back by the same amount the direction has been streched
    lengthOfProjectedDirection := projectedDirection.Length()
    normalizedProjectedDirection := vec3.Divide(projectedDirection, lengthOfProjectedDirection)
    
    projectedRay := ray.Ray{
        projectedOrigin,
        normalizedProjectedDirection,
        r.T0,
        r.T1,
    }
    
    projectedHit := this.collisionSphere.Intersect(projectedRay)
    
    if projectedHit == nil {
        return nil
    }
    
    // Scale T back to direction-pre-normalized state
    scaledT := projectedHit.T / lengthOfProjectedDirection
    
    point := r.PointAt(scaledT)
    
    positionToPoint := vec3.Subtract(point, this.Position)
    // "beneath" the cylinder
    if vec3.DotProduct(positionToPoint, this.Normal) < 0 {
        return nil
    }
    positionToPointProjectedOnNormalSquaredLength := positionToPoint.SquaredLength() - this.Radius*this.Radius
    // "above" the cylinder
    if positionToPointProjectedOnNormalSquaredLength > this.Height*this.Height {
        return nil
    }
    
    return &space.Hit{
        T: scaledT,
        Position: point,
        Normal: projectedHit.Normal,
        Material: this.Material,
        // Material: Material_Diffuse{
        //     projectedHit.Normal,
        // },
    }
}










