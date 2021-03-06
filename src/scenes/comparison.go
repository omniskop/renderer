package scenes

import (
	"math"
	"renderer/camera"
	"renderer/shapes"
	"space"
	"space/texture"
	"vec3"
)

func GetComparisonScene() shapes.Group {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			/*shapes.Sphere{
			      vec3.Vec3{-2,0,-30},
			      1,
			      space.Material_Diffuse{vec3.Vec3{0.7,1,0.7}},
			  },
			  shapes.Sphere{
			      vec3.Vec3{0,0,-30},
			      1,
			      space.Material_Metal{vec3.Vec3{1,0.7,0.7}, 0},
			  },
			  shapes.Sphere{
			      vec3.Vec3{2,0,-30},
			      1,
			      space.Material_Diffuse{vec3.Vec3{0.7,0.7,1}},
			  },*/
			shapes.Sphere{
				vec3.Vec3{0, -0.5, -15},
				1,
				space.Material_Transparent{texture.NewColor(1, 1, 1), 1.3},
			},
			shapes.Sphere{
				vec3.Vec3{-0.9, -0.9, -14},
				0.2,
				space.Material_Diffuse{texture.NewColor(1, 0, 0)},
			},
			shapes.Plane{vec3.Vec3{0, -1, 0}, vec3.Vec3{0, 1, 0}, space.Material_Metal{vec3.Vec3{0.1, 0.1, 0.4}, 0.05}},
			shapes.Plane{vec3.Vec3{0, -1.1, 0}, vec3.Vec3{0, 1, 0}, space.Material_Diffuse{texture.NewColor(1, 0, 0)}},
			shapes.Background{space.Material_Sky{texture.NewColor(1, 1, 1)}},
		},
	}
}

func GetComparisonSceneCamera() camera.Camera {
	return camera.PinholeCamera{
		Position:     vec3.Vec3{0, 0, 0},
		Direction:    vec3.Vec3{0, 0, -1},
		OpeningAngle: math.Pi / 10,
		Width:        200,
		Height:       200,
	}
}
