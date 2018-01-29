package scenes

import (
	"math"
	"renderer/camera"
	"renderer/shapes"
	"space"
	"space/texture"
	"vec3"
)

func GetTextureScene() shapes.Group {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Sphere{vec3.Vec3{0, 1, 0}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},
			shapes.Sphere{vec3.Vec3{2, 1, 2}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},
			shapes.Sphere{vec3.Vec3{0, 1, 4}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},
			shapes.Sphere{vec3.Vec3{-2, 1, 2}, 1, space.Material_Diffuse{texture.NewTexturemap("res/tag.png", 2.2)}},
			// shapes.Plane{vec3.Vec3{0,0,0}, vec3.Vec3{0,1,0}, space.Material_Metal{vec3.Vec3{0.9,0.9,1}, 0.05}},
			shapes.Plane{vec3.Vec3{0, 0, 0}, vec3.Vec3{0, 1, 0}, space.Material_Diffuse{texture.NewScaledTexturemap("res/gravel.jpg", 2.2, 0.3)}},
			shapes.Background{space.Material_Sky{texture.NewTexturemap("res/sky2.jpg", 2.2)}},
			// shapes.Background{space.Material_Sky{texture.NewColor(vec3.Vec3{1, 1, 1})}},
		},
	}
}

func GetTextureSceneCamera() camera.Camera {
	return camera.SphereCamera{
		Position:  vec3.Vec3{0, 1, 2},
		Direction: vec3.Normalize(vec3.Vec3{0, 0, -0.5}),
		Width:     2000,
		Height:    1000,
	}
}

func GetTextureSceneCamera2() camera.Camera {
	return camera.PinholeCamera{
		Position:     vec3.Vec3{0, 3, 10},
		Direction:    vec3.Normalize(vec3.Vec3{0, -0.1, -0.5}),
		OpeningAngle: math.Pi / 8,
		Width:        200,
		Height:       200,
	}
}
