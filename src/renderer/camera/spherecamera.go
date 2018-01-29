package camera

import (
	"mat4"
	"math"
	"space/ray"
	"vec3"
)

type SphereCamera struct {
	Position  vec3.Vec3
	Direction vec3.Vec3
	Tilt      float64
	// transformationMatrix    mat4.mat4
	Width  int
	Height int
}

func (this SphereCamera) GetWidth() int           { return this.Width }
func (this SphereCamera) GetHeight() int          { return this.Height }
func (this SphereCamera) GetPosition() vec3.Vec3  { return this.Position }
func (this SphereCamera) GetDirection() vec3.Vec3 { return this.Direction }

func (this SphereCamera) GetRayForPixel(x float64, y float64) ray.Ray {

	polarwinkel := (y / float64(this.Height)) * math.Pi
	azimuth := (x / float64(this.Width)) * math.Pi * 2

	dir := vec3.Normalize(vec3.Vec3{
		math.Sin(polarwinkel) * math.Sin(azimuth),
		math.Cos(polarwinkel),
		math.Sin(polarwinkel) * math.Cos(azimuth),
	})

	transformationMatrix := mat4.NewTranslationByVec3(this.Position).Multiply(
		mat4.NewRotationAroundVec3(vec3.Vec3{0, 1, 0}, -math.Asin(this.Direction.X)),
	)

	cameraPosition := transformationMatrix.TransformPoint(vec3.Vec3{0, 0, 0})
	cameraDirection := transformationMatrix.TransformDirection(dir)

	out := ray.Ray{
		Origin:    cameraPosition,
		Direction: cameraDirection,
		T0:        0,
		T1:        math.Inf(1),
	}

	out.Direction.Normalize()

	return out
}
