package shapes

import (
	"math"
	"space"
	"space/ray"
	"vec3"
)

// Cube represents a cube
type Cube struct {
	Position vec3.Vec3
	Size     vec3.Vec3
	Material space.Material
}

// Intersect returns the first hit of a ray with the object
func (cube Cube) Intersect(r ray.Ray) *space.Hit {
	r2 := r
	max := vec3.Add(cube.Position, cube.Size)
	tnear := math.Inf(-1)
	tfar := math.Inf(1)

	{
		t1 := (cube.Position.X - r2.Origin.X) / r2.Direction.X
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
		t1 := (cube.Position.Y - r2.Origin.Y) / r2.Direction.Y
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
		t1 := (cube.Position.Z - r2.Origin.Z) / r2.Direction.Z
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

	if math.Abs(point.X-cube.Position.X) < epislon {
		normal = vec3.Vec3{-1, 0, 0}
		surfaceCoordinates = vec3.Vec3{(point.Z - cube.Position.Z) / cube.Size.Z, -(point.Y - cube.Position.Y) / cube.Size.Y, 1}
	} else if math.Abs(point.X-max.X) < epislon {
		normal = vec3.Vec3{1, 0, 0}
		surfaceCoordinates = vec3.Vec3{-(point.Z - cube.Position.Z) / cube.Size.Z, -(point.Y - cube.Position.Y) / cube.Size.Y, 1}
	} else {
		if math.Abs(point.Y-cube.Position.Y) < epislon {
			normal = vec3.Vec3{0, -1, 0}
			surfaceCoordinates = vec3.Vec3{(point.X - cube.Position.X) / cube.Size.X, (point.Z - cube.Position.Z) / cube.Size.Z, 0}
		} else if math.Abs(point.Y-max.Y) < epislon {
			normal = vec3.Vec3{0, 1, 0}
			surfaceCoordinates = vec3.Vec3{-(point.X - cube.Position.X) / cube.Size.X, (point.Z - cube.Position.Z) / cube.Size.Z, 2}
		} else {
			if math.Abs(point.Z-cube.Position.Z) < epislon {
				normal = vec3.Vec3{0, 0, -1}
				surfaceCoordinates = vec3.Vec3{(point.X - cube.Position.X) / cube.Size.X, (point.Y - cube.Position.Y) / cube.Size.Y, 1}
			} else if math.Abs(point.Z-max.Z) < epislon {
				normal = vec3.Vec3{0, 0, 1}
				surfaceCoordinates = vec3.Vec3{(point.X - cube.Position.X) / cube.Size.X, -(point.Y - cube.Position.Y) / cube.Size.Y, 1}
			} else {
				panic("Unable to determine normal vector of cube. Try to increase epsilon.")
			}
		}
	}
	return &space.Hit{
		T:                  tnear,
		Position:           point,
		Normal:             normal,
		SurfaceCoordinates: surfaceCoordinates,
		Material:           cube.Material,
		IgnoreLight:        cube.Size.X == 0.1,
	}
}

// Includes checks if the point is inside the object
func (cube Cube) Includes(point vec3.Vec3) bool {
	p := vec3.Add(cube.Position, vec3.Multiply(.5, cube.Size))
	if math.Abs(math.Max(math.Max(math.Abs(point.X-p.X), math.Abs(point.Y-p.Y)), math.Abs(point.Z-p.Z))-cube.Size.X) < 0.00001 {
		return true
	}
	return false
}
