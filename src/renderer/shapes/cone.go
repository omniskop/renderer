package shapes

import (
	"log"
	"math"
	"space"
	"space/ray"
	"vec3"
)

type Cone struct {
	Position vec3.Vec3
	Normal   vec3.Vec3
	Radius   float64
	Height   float64
	Material space.Material
	disc     Disc
	walls    OpenCone
}

func NewCone(position, normal vec3.Vec3, radius, height float64, material space.Material) *Cone {
	out := Cone{Position: position, Normal: normal, Radius: radius / height, Height: height, Material: material}
	// out.disc = Disc{position, vec3.Vec3{0,-1,0}, radius, material}
	out.disc = Disc{position, vec3.Multiply(-1, normal), radius, material}
	out.walls = NewOpenCone(position, normal, radius, height, material)

	return &out
}

func (this Cone) Intersect(r ray.Ray) *space.Hit {
	if vec3.DotProduct(this.disc.Normal, r.Direction) < 0 {
		discHit := this.disc.Intersect(r)

		if discHit != nil {
			return discHit
		}
	}

	coneHit := this.walls.Intersect(r)

	return coneHit
}

type OpenCone struct {
	Position vec3.Vec3
	Normal   vec3.Vec3
	Radius   float64
	Height   float64
	Material space.Material
}

func NewOpenCone(position, normal vec3.Vec3, radius, height float64, material space.Material) OpenCone {
	out := OpenCone{Position: position, Normal: normal, Radius: radius / height, Height: height, Material: material}
	return out
}

func (this OpenCone) Intersect(r ray.Ray) *space.Hit {
	d := r.Direction
	// x0 :=  vec3.Subtract(vec3.Subtract(r.Origin, this.Position), vec3.Vec3{0,this.Height,0})
	x0 := vec3.Subtract(vec3.Subtract(r.Origin, this.Position), vec3.Multiply(this.Height, this.Normal))

	a := d.X*d.X + d.Z*d.Z - (d.Y*d.Y)*(this.Radius*this.Radius)
	b := 2*x0.X*d.X + 2*x0.Z*d.Z - (2*x0.Y*d.Y)*(this.Radius*this.Radius)
	c := x0.X*x0.X + x0.Z*x0.Z - (x0.Y*x0.Y)*(this.Radius*this.Radius)

	sqrtOf := b*b - 4*a*c

	_ = log.Print

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

	if point.Y < this.Position.Y || point.Y > this.Position.Y+this.Height {
		return nil
	}

	// tipToPoint := vec3.Subtract(point,vec3.Add(this.Position, vec3.Vec3{0,this.Height,0}))
	tipToPoint := vec3.Subtract(point, vec3.Add(this.Position, vec3.Multiply(this.Height, this.Normal)))
	// normal := vec3.Normalize(vec3.CrossProduct(vec3.CrossProduct(tipToPoint, vec3.Vec3{0,1,0}), tipToPoint))
	normal := vec3.Normalize(vec3.CrossProduct(vec3.CrossProduct(tipToPoint, this.Normal), tipToPoint))

	// it is correct. Don't change it again!
	if vec3.DotProduct(r.Direction, normal) > 0 {
		normal = vec3.Multiply(-1, normal)
	}

	return &space.Hit{
		T:        t,
		Position: point,
		Normal:   normal,
		Material: this.Material,
		// Material: Material_Normal{},
	}
}

func (this Cone) Includes(point vec3.Vec3) bool {
	if point.Y < this.Position.Y || point.Y > this.Position.Y+this.Height {
		return false
	}
	if vec3.Subtract(vec3.Vec3{point.X, 0, point.Z}, this.Position).SquaredLength() > (point.Y-this.Position.Y)/this.Height*this.Radius {
		return false
	}
	return true
}

func (this OpenCone) Includes(point vec3.Vec3) bool {
	return false
}
