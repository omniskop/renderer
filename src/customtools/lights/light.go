package lights

import (
	"customtools/shapes"
	"customtools/space"
	"customtools/vec3"
	"math"
)

type Light interface {
	Sample(world shapes.Shape, point vec3.Vec3) (vec3.Vec3, vec3.Vec3)
}

type Ambient struct {
	Emission vec3.Vec3
}

func (this Ambient) Sample(world shapes.Shape, point vec3.Vec3) (vec3.Vec3, vec3.Vec3) {
	return vec3.Zero, this.Emission
}

func SampleLights(world shapes.Shape, light []Light, objectHit *space.Hit) (emi vec3.Vec3) {
	if math.IsInf(objectHit.T, 0) { // Filter out background hits
		return vec3.Zero
	}
	if len(light) == 0 {
		return vec3.Zero
	}
	_, emi = light[0].Sample(world, objectHit.Position)
	for i := 1; i < len(light); i++ {
		direction, emission := light[i].Sample(world, objectHit.Position)
		if direction.Equals(vec3.Zero) {
			continue
		}
		angle := vec3.DotProduct(direction, objectHit.Normal)
		if angle <= 0 {
			continue
		}
		emi = vec3.Add(emi, vec3.Multiply(angle, emission))
	}
	return emi
}
