package lights

import (
	"renderer/shapes"
	"space/ray"
	"vec3"
)

type Point struct {
	Position vec3.Vec3
	Emission vec3.Vec3
}

func (this Point) Sample(world shapes.Shape, point vec3.Vec3) (vec3.Vec3, vec3.Vec3) {
	vectorToLight := vec3.Subtract(this.Position, point)
	normalToLight := vec3.Normalize(vectorToLight)
	hit := world.Intersect(ray.Ray{
		point,
		normalToLight,
		0.000001,
		vectorToLight.Length(),
	})
	if hit == nil {
		return normalToLight, this.Emission
	}
	return vec3.Zero, vec3.Black
}
