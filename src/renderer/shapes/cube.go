package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

type Cube struct {
	Position vec3.Vec3
	Size     vec3.Vec3
	Material space.Material
}

func (this Cube) Intersect(r ray.Ray) *space.Hit {
	r2 := r
	max := vec3.Add(this.Position, this.Size)
	tnear := math.Inf(-1)
	tfar := math.Inf(1)

	{
		t1 := (this.Position.X - r2.Origin.X) / r2.Direction.X
		t2 := (max.X - r2.Origin.X) / r2.Direction.X
		if t1 > t2 {
			t1, t2 = t2, t1
		}
		if t1 > tnear {
			tnear = t1
		}
		if t2 < tfar {
			tfar = t2
		}
		if tnear > tfar {
			return nil
		}
		if tfar < 0 {
			return nil
		}
	}

	{
		t1 := (this.Position.Y - r2.Origin.Y) / r2.Direction.Y
		t2 := (max.Y - r2.Origin.Y) / r2.Direction.Y
		if t1 > t2 {
			t1, t2 = t2, t1
		}
		if t1 > tnear {
			tnear = t1
		}
		if t2 < tfar {
			tfar = t2
		}
		if tnear > tfar {
			return nil
		}
		if tfar < 0 {
			return nil
		}
	}

	{
		t1 := (this.Position.Z - r2.Origin.Z) / r2.Direction.Z
		t2 := (max.Z - r2.Origin.Z) / r2.Direction.Z
		if t1 > t2 {
			t1, t2 = t2, t1
		}
		if t1 > tnear {
			tnear = t1
		}
		if t2 < tfar {
			tfar = t2
		}
		if tnear > tfar {
			return nil
		}
		if tfar < 0 {
			return nil
		}
	}

	point := r2.PointAt(tnear)
	const epislon = 0.00001
	var normal vec3.Vec3
	var surfaceCoordinates vec3.Vec3

	if math.Abs(point.X-this.Position.X) < epislon {
		normal = vec3.Vec3{-1, 0, 0}
		surfaceCoordinates = vec3.Vec3{(point.Z - this.Position.Z) / this.Size.Z, -(point.Y - this.Position.Y) / this.Size.Y, 1}
	} else if math.Abs(point.X-max.X) < epislon {
		normal = vec3.Vec3{1, 0, 0}
		surfaceCoordinates = vec3.Vec3{-(point.Z - this.Position.Z) / this.Size.Z, -(point.Y - this.Position.Y) / this.Size.Y, 1}
	} else {
		if math.Abs(point.Y-this.Position.Y) < epislon {
			normal = vec3.Vec3{0, -1, 0}
			surfaceCoordinates = vec3.Vec3{(point.X - this.Position.X) / this.Size.X, (point.Z - this.Position.Z) / this.Size.Z, 0}
		} else if math.Abs(point.Y-max.Y) < epislon {
			normal = vec3.Vec3{0, 1, 0}
			surfaceCoordinates = vec3.Vec3{-(point.X - this.Position.X) / this.Size.X, (point.Z - this.Position.Z) / this.Size.Z, 2}
		} else {
			if math.Abs(point.Z-this.Position.Z) < epislon {
				normal = vec3.Vec3{0, 0, -1}
				surfaceCoordinates = vec3.Vec3{(point.X - this.Position.X) / this.Size.X, (point.Y - this.Position.Y) / this.Size.Y, 1}
			} else if math.Abs(point.Z-max.Z) < epislon {
				normal = vec3.Vec3{0, 0, 1}
				surfaceCoordinates = vec3.Vec3{(point.X - this.Position.X) / this.Size.X, -(point.Y - this.Position.Y) / this.Size.Y, 1}
			} else {
				panic("Unable to determine normal vector of cube. Try to increase epsilon.")
			}
		}
	}

	// r.Direction = vec3.MultiplyByVec3(r.Direction, this.Size)
	// r.Origin = vec3.MultiplyByVec3(r.Direction, this.Size)
	// point = r.PointAt(tnear)

	return &space.Hit{
		T:                  tnear,
		Position:           point,
		Normal:             normal,
		SurfaceCoordinates: surfaceCoordinates,
		Material:           this.Material,
		IgnoreLight:        this.Size.X == 0.1,
	}
}

func (this Cube) Includes(point vec3.Vec3) bool {
	p := vec3.Add(this.Position, vec3.Multiply(.5, this.Size))
	if math.Abs(math.Max(math.Max(math.Abs(point.X-p.X), math.Abs(point.Y-p.Y)), math.Abs(point.Z-p.Z))-this.Size.X) < 0.00001 {
		return true
	}
	return false

	// if point.X > this.Position && point.X < this.Position
}
