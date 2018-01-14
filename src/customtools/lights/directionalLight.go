package lights

import (
	"customtools/ray"
	"customtools/shapes"
	"customtools/vec3"
	"math"
)

type Directional struct {
	Direction vec3.Vec3
	Emission  vec3.Vec3
}

func (this Directional) Sample(world shapes.Shape, point vec3.Vec3) (vec3.Vec3, vec3.Vec3) {
	hit := world.Intersect(ray.Ray{
		point,
		vec3.Multiply(-1, this.Direction),
		0.00001,
		math.Inf(1),
	})
	// Hit can't be null because there will always be a background. If not the main code will already throw an error.
	if math.IsInf(hit.T, 0) {
		return vec3.Multiply(-1, this.Direction), this.Emission
	}
	return vec3.Zero, vec3.Black
}
