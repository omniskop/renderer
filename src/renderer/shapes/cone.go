package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// Cone represents a cone
type Cone struct {
	Position vec3.Vec3
	Normal   vec3.Vec3
	Radius   float64
	Height   float64
	Material space.Material
	disc     Disc
	walls    OpenCone
}

// NewCone creates and returns a new cone
func NewCone(position, normal vec3.Vec3, radius, height float64, material space.Material) *Cone {
	out := Cone{Position: position, Normal: normal, Radius: radius / height, Height: height, Material: material}
	// out.disc = Disc{position, vec3.Vec3{0,-1,0}, radius, material}
	out.disc = Disc{position, vec3.Multiply(-1, normal), radius, material}
	out.walls = NewOpenCone(position, normal, radius, height, material)

	return &out
}

// Intersect returns the first hit of a ray with the object
func (cone Cone) Intersect(r ray.Ray) *space.Hit {
	if vec3.DotProduct(cone.disc.Normal, r.Direction) < 0 {
		discHit := cone.disc.Intersect(r)

		if discHit != nil {
			return discHit
		}
	}

	coneHit := cone.walls.Intersect(r)

	return coneHit
}

// OpenCone represents a cone without a base plane
type OpenCone struct {
	Position vec3.Vec3
	Normal   vec3.Vec3
	Radius   float64
	Height   float64
	Material space.Material
}

// NewOpenCone creates and returns a new open cone
func NewOpenCone(position, normal vec3.Vec3, radius, height float64, material space.Material) OpenCone {
	out := OpenCone{Position: position, Normal: normal, Radius: radius / height, Height: height, Material: material}
	return out
}

// Intersect returns the first hit of a ray with the object
func (oCone OpenCone) Intersect(r ray.Ray) *space.Hit {
	d := r.Direction
	// x0 :=  vec3.Subtract(vec3.Subtract(r.Origin, this.Position), vec3.Vec3{0,this.Height,0})
	x0 := vec3.Subtract(vec3.Subtract(r.Origin, oCone.Position), vec3.Multiply(oCone.Height, oCone.Normal))

	a := d.X*d.X + d.Z*d.Z - (d.Y*d.Y)*(oCone.Radius*oCone.Radius)
	b := 2*x0.X*d.X + 2*x0.Z*d.Z - (2*x0.Y*d.Y)*(oCone.Radius*oCone.Radius)
	c := x0.X*x0.X + x0.Z*x0.Z - (x0.Y*x0.Y)*(oCone.Radius*oCone.Radius)

	sqrtOf := b*b - 4*a*c

	if sqrtOf < 0 {
		return nil
	}
	var t float64
	if sqrtOf == 0 {
		t = -b / (2 * a)
		if !r.InRange(t) {
			return nil
		}
	} else if sqrtOf > 0 {
		t = (-b - math.Sqrt(sqrtOf)) / (2 * a)
		if !r.InRange(t) {
			t = (-b + math.Sqrt(sqrtOf)) / (2 * a)
			if !r.InRange(t) {

				return nil
			}
		}
	}

	point := r.PointAt(t)

	if point.Y < oCone.Position.Y || point.Y > oCone.Position.Y+oCone.Height {
		return nil
	}

	tipToPoint := vec3.Subtract(point, vec3.Add(oCone.Position, vec3.Multiply(oCone.Height, oCone.Normal)))
	normal := vec3.Normalize(vec3.CrossProduct(vec3.CrossProduct(tipToPoint, oCone.Normal), tipToPoint))

	// it is correct. Don't change it again!
	if vec3.DotProduct(r.Direction, normal) > 0 {
		normal = vec3.Multiply(-1, normal)
	}

	return &space.Hit{
		T:        t,
		Position: point,
		Normal:   normal,
		Material: oCone.Material,
		// Material: Material_Normal{},
	}
}

// Includes checks if the point is inside the object
func (cone Cone) Includes(point vec3.Vec3) bool {
	if point.Y < cone.Position.Y || point.Y > cone.Position.Y+cone.Height {
		return false
	}
	if vec3.Subtract(vec3.Vec3{point.X, 0, point.Z}, cone.Position).SquaredLength() > (point.Y-cone.Position.Y)/cone.Height*cone.Radius {
		return false
	}
	return true
}

// Includes checks if the point is inside the object
func (oCone OpenCone) Includes(point vec3.Vec3) bool {
	return false
}
