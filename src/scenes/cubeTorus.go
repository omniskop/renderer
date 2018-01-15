package scenes

import (
	"customtools/camera"
	"customtools/lights"
	"customtools/shapes"
	"customtools/space"
	"customtools/texture"
	"customtools/vec3"
	"math"
)

func GetCubeTorusScene() shapes.Group {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Sphere{vec3.Vec3{0, 0, 0}, 0.5, space.Material_Diffuse{texture.NewColor(1, 0, 0)}},

			shapes.Cube{vec3.Vec3{0, 0, -1}, vec3.Vec3{5, 5, 5}, space.Material_Diffuse{texture.NewTexturemap("res/wood.jpg", 1)}},
			shapes.Cube{vec3.Vec3{1.5, 0, -1 + 5}, vec3.Vec3{2, 2, 2}, space.Material_Diffuse{texture.NewTexturemap("res/wood.jpg", 1)}},
			shapes.Cube{vec3.Vec3{5, 0, -1 + 1.5}, vec3.Vec3{2, 2, 2}, space.Material_Diffuse{texture.NewTexturemap("res/wood.jpg", 1)}},

			shapes.Cube{vec3.Vec3{3.5, 0, -1 + 6}, vec3.Vec3{2, 1, 1}, space.Material_Diffuse{texture.NewTexturemap("res/wood.jpg", 1)}},
			shapes.Cube{vec3.Vec3{6, 0, -1 + 1.5 + 2}, vec3.Vec3{1, 1, 2}, space.Material_Diffuse{texture.NewTexturemap("res/wood.jpg", 1)}},

			shapes.Plane{vec3.Vec3{0, 0, 0}, vec3.Vec3{0, 1, 0}, space.Material_Diffuse{texture.NewScaledTexturemap("res/gravel.jpg", 2.2, 0.3)}},
			shapes.Background{space.Material_Sky{texture.NewTexturemap("res/sky2.jpg", 2.2)}},
		},
	}
}

func GetCubeTorusSceneLight() []lights.Light {
	return []lights.Light{
		lights.Ambient{
			vec3.Vec3{0.1, 0.1, 0.1},
		},
		lights.Point{
			vec3.Vec3{7, 7, 7 - 1},
			vec3.Vec3{2, 2, 2},
		},
	}
}

func GetCubeTorusSceneCamera() camera.Camera {
	return camera.PinholeCamera{
		Position:     vec3.Vec3{10, 10, 10},
		Direction:    vec3.Normalize(vec3.Vec3{-0.5, -0.5, -0.5}),
		OpeningAngle: math.Pi / 8,
		Width:        200,
		Height:       200,
	}
}

func GetCubeTorusSceneCamera2() camera.Camera {
	return camera.PinholeCamera{
		Position:     vec3.Vec3{15, 13, 15},
		Direction:    vec3.Normalize(vec3.Vec3{-0.6, -0.4, -0.5}),
		OpeningAngle: math.Pi / 8,
		Width:        200,
		Height:       200,
	}
}
