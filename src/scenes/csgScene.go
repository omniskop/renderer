package scenes

import (
	"customtools/camera"
	"customtools/shapes"
	"customtools/space"
	"customtools/texture"
	"customtools/vec3"
	"math"
)

func GetCsgScene() shapes.Group {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Union( // main sphere union
				shapes.Sphere{
					vec3.Vec3{-0.5, 0, 0},
					1,
					space.Material_Diffuse{texture.NewColor(1, 0, 0)},
				},
				shapes.Sphere{
					vec3.Vec3{0.5, 0, 0},
					1,
					space.Material_Diffuse{texture.NewColor(0, 0, 1)},
				},
			),
			shapes.Difference( // main sphere difference
				shapes.Sphere{
					vec3.Vec3{-0.5, 3, 0},
					1,
					space.Material_Diffuse{texture.NewColor(1, 0, 0)},
				},
				shapes.Sphere{
					vec3.Vec3{0.5, 3, 0},
					1,
					space.Material_Diffuse{texture.NewColor(0, 0, 1)},
				},
			),
			shapes.Group{
				space.NewTransformation(vec3.Vec3{2.5, 0, -1}, 0, math.Pi/15, 0),
				[]shapes.Shape{shapes.Difference( // difference sphere rotated
					shapes.Sphere{
						vec3.Vec3{-0.5, 3, 0},
						1,
						space.Material_Diffuse{texture.NewColor(1, 0, 0)},
					},
					shapes.Sphere{
						vec3.Vec3{0.5, 3, 0},
						1,
						space.Material_Diffuse{texture.NewColor(0, 0, 1)},
					},
				),
					shapes.Disc{
						vec3.Vec3{-0.5, 3, 0},
						vec3.Vec3{0, 1, 0},
						1.5,
						space.Material_Diffuse{texture.NewColor(0, 1, 0)},
					}},
			},
			shapes.Difference( // difference cylinder
				shapes.NewCylinder(
					vec3.Vec3{-0.5 + 1 + 0.5, 3, 0},
					vec3.Vec3{1, 0, 0},
					0.5,
					2,
					space.Material_Diffuse{texture.NewColor(1, 0, 0)},
				),
				shapes.NewCylinder(
					vec3.Vec3{-0.5 + 1 + 0.5 + 1, 3 - 1, 0},
					vec3.Vec3{0, 1, 0},
					0.5,
					2,
					space.Material_Diffuse{texture.NewColor(0, 0, 1)},
				),
			),
			shapes.Intersection( // main sphere intersection
				shapes.Sphere{
					vec3.Vec3{-0.5, 6, 0},
					1,
					space.Material_Diffuse{texture.NewColor(1, 0, 0)},
				},
				shapes.Sphere{
					vec3.Vec3{0.5, 6, 0},
					1,
					space.Material_Diffuse{texture.NewColor(0, 0, 1)},
				},
			),
			shapes.Intersection( // main sphere intersection
				shapes.NewCone(
					vec3.Vec3{-2, 6, 0},
					vec3.Vec3{0, 1, 0},
					1,
					1,
					space.Material_Diffuse{texture.NewColor(1, 0, 0)},
				),
				shapes.Group{
					space.NewTransformation(vec3.Vec3{0, 0, 0}, math.Pi, 0, 0),
					[]shapes.Shape{
						shapes.NewCone(
							vec3.Vec3{-2, -6.5, 0},
							vec3.Vec3{0, 1, 0},
							1,
							1,
							space.Material_Diffuse{texture.NewColor(0, 0, 1)},
						),
					},
				},
			),
			shapes.Disc{vec3.Vec3{-2, 6, 0}, vec3.Vec3{0, 0, 1}, 1, space.Material_Diffuse{texture.NewColor(0, 1, 0)}},
			// shapes.Group{
			// 	space.NewTransformation(vec3.Vec3{0, 0, 0}, math.Pi, 0, 0),
			// 	[]shapes.Shape{
			// 		shapes.NewCone(
			// 			vec3.Vec3{-2, -6.5, 0},
			// 			vec3.Vec3{0, 1, 0},
			// 			1,
			// 			1,
			// 			space.Material_Diffuse{texture.NewColor(0, 0, 1)},
			// 		),
			// 		// shapes.Disc{vec3.Vec3{0, -6, 0}, vec3.Vec3{0, -1, 0}, 1, space.Material_Diffuse{texture.NewColor(0, 1, 0)}},
			// 	},
			// },
			// shapes.Plane{vec3.Vec3{0, 3, 0}, vec3.Vec3{0, 1, 0}, space.Material_Diffuse{texture.NewColor(0, 1, 0)}},
			shapes.Background{space.Material_Sky{texture.NewColor(1, 1, 1)}},
		},
	}
}

func GetCsgSceneCamera() camera.Camera {
	// return camera.PinholeCamera{
	// 	Position:     vec3.Vec3{0, 3, 10},
	// 	Direction:    vec3.Normalize(vec3.Vec3{0, 0, -1}),
	// 	OpeningAngle: math.Pi / 2,
	// 	Width:        400,
	// 	Height:       400,
	// }
	return camera.PinholeCamera{
		Position:     vec3.Vec3{0, 8, 10},
		Direction:    vec3.Normalize(vec3.Vec3{0, -0.5, -1}),
		OpeningAngle: math.Pi / 2,
		Width:        400,
		Height:       400,
	}
}
