package scenes

import (
	"math"
	"renderer/camera"
	"renderer/lights"
	"renderer/shapes"
	"space"
	"space/texture"
	"vec3"
)

func GetLightScene() shapes.Group {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Sphere{vec3.Vec3{0, 1, -2}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},
			shapes.Sphere{vec3.Vec3{2, 1, 0}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},
			// shapes.Sphere{vec3.Vec3{2, 1, 0}, 1, space.Material_Diffuse{texture.NewColor(1, 0, 0)}},
			shapes.Sphere{vec3.Vec3{0, 1, 2}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},
			shapes.Sphere{vec3.Vec3{-2, 1, 0}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},

			shapes.Disc{vec3.Vec3{0, 5, 0}, vec3.Vec3{0, 1, 0}, 5, space.Material_Diffuse{texture.NewColor(1, 0, 0)}},
			shapes.Sphere{vec3.Vec3{0, 7, 0}, 1, space.Material_Diffuse{texture.NewColor(1, 1, 1)}},

			shapes.Plane{vec3.Vec3{0, 0, 0}, vec3.Vec3{0, 1, 0}, space.Material_Diffuse{texture.NewScaledTexturemap("res/gravel.jpg", 2.2, 0.3)}},
			shapes.Background{space.Material_Sky{texture.NewTexturemap("res/sky2.jpg", 2.2)}},
		},
	}
}

func GetLightSceneLight() []lights.Light {
	return []lights.Light{
		lights.Ambient{
			vec3.Vec3{0.1, 0.1, 0.1},
		},
		lights.Point{
			vec3.Vec3{0, 1, 0},
			vec3.Vec3{2, 2, 0},
		},
		lights.Directional{
			vec3.Normalize(vec3.Vec3{0, -1, 0}),
			vec3.Vec3{0.2, 0.2, 2},
		},
	}
}

func GetLightSceneCamera1() camera.Camera {
	return camera.PinholeCamera{
		Position:     vec3.Vec3{0, 6, 20},
		Direction:    vec3.Normalize(vec3.Vec3{0, -0.1, -0.5}),
		OpeningAngle: math.Pi / 8,
		Width:        200,
		Height:       200,
	}
}

func GetLightSceneCamera() camera.Camera {
	return camera.PinholeCamera{
		Position:     vec3.Vec3{15, 13, 15},
		Direction:    vec3.Normalize(vec3.Vec3{-0.6, -0.4, -0.5}),
		OpeningAngle: math.Pi / 8,
		Width:        200,
		Height:       200,
	}
}
