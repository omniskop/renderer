package scenes

import (
	"renderer/shapes"
	"space"
	"space/texture"
	"vec3"
)

/*
Position: vec3.Vec3{0,0,4},
Direction: vec3.Normalize(vec3.Vec3{0,-0.01,-1}),
OpeningAngle: math.Pi / 7,
*/
func GetWaterMolecule() shapes.Group {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Sphere{ // Red sphere
				Position: vec3.Vec3{0, -0.4, -11.2},
				Material: space.Material_Diffuse{texture.NewColor(1, 0, 0)},
				Radius:   1,
			},
			shapes.Sphere{ // left blue sphere
				Position: vec3.Vec3{-.7, -1, -10.7},
				Material: space.Material_Diffuse{texture.NewColor(0, 0, 1)},
				Radius:   .5,
			},
			shapes.Sphere{ // right blue sphere
				Position: vec3.Vec3{.7, -1, -10.7},
				Material: space.Material_Diffuse{texture.NewColor(0, 0, 1)},
				Radius:   .5,
			},
			shapes.Sphere{ // mirror sphere
				Position: vec3.Vec3{0, 5.7, -11.2},
				Material: space.Material_Metal{vec3.White, 0.05},
				Radius:   5,
			},
			shapes.Sphere{
				Position: vec3.Vec3{0.5, 0.5, -6},
				Material: space.Material_Transparent{texture.NewColor(1, 1, 1), 1.5},
				Radius:   2,
			},
			shapes.Plane{ // bottom
				Position: vec3.Vec3{0, -1.3, 0},
				Normal:   vec3.Vec3{0, 1, 0},
				// Material: space.Material_Diffuse_Checkerboard{1,vec3.White, vec3.Black},
				Material: space.Material_Metal_Checkerboard{1, vec3.White, vec3.Vec3{0.01, 0.01, 0.01}, 0.005},
			},
			shapes.Background{space.Material_Sky{texture.NewColor(0.8, 0.8, 1)}},
		},
	}
}
