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

func GetMinecraftScene() shapes.Group {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			// tree(0, 1, 0),
			layer0(),
			layer1(),
			layer2(),
			layer3(),
			// shapes.Sphere{vec3.Vec3{0, 0, 0}, 1, space.Material_Diffuse{texture.NewColor(1, 0, 0)}},
			shapes.Plane{vec3.Vec3{0, -0.2, 0}, vec3.Vec3{0, 1, 0}, space.Material_Diffuse{texture.NewScaledTexturemap("res/mc/water.jpg", 2.2, 0.3)}},
			shapes.Background{space.Material_Sky{texture.NewTexturemap("res/sky2.jpg", 2.2)}},
		},
	}
}

func layer3() shapes.Shape {
	var s float64 = 2.5
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			tree(-5*s, 0, 3*s),
			tree(-4*s, 0, -1*s),
			tree(-4*s, 0, -4*s),
			tree(-1*s, 0, -7*s),
			tree(3*s, 0, -8*s),
			tree(6*s, 0, -7*s),
			tree(8*s, 0, -5*s),
			tree(8*s, 0, -2*s),

			dirtArea(-27+20-3, 4, -10-20-7, 25),
		},
	}
}

func layer2() shapes.Shape {
	var s float64 = 2
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			torch(-2.5, 1, 2),

			grassArea(-27, 1, -10, 20),
			grassArea(-27+3, 1, -28, 20),
			grassArea(-27+20, 1, -10-20, 20),
			grassArea(-27+37, 1, -10-20+10, 20),

			tree(-5*s, 1, 3*s),
			tree(-4*s, 1, -1*s),
			tree(-4*s, 1, -4*s),
			tree(-1*s, 1, -7*s),
			tree(3*s, 1, -8*s),
			tree(6*s, 1, -7*s),
			tree(8*s, 1, -5*s),
			tree(8*s, 1, -2*s),

			grass(-4*s, 1, 0*s),
			grass(-4*s, 1, -1*s),
			grass(-4*s, 1, -2*s),
			grass(-4*s, 1, -3*s),
			grass(-4*s, 1, -4*s),
			grass(-4*s, 1, -5*s), grass(-4*s, 1, -6*s), grass(-4*s, 1, -7*s),
			grass(-3*s, 1, -4*s),
			grass(-3*s, 1, -5*s),
			grass(-3*s, 1, -6*s), grass(-3*s, 1, -7*s), grass(-2*s, 1, -7*s), grass(-1*s, 1, -7*s),
			grass(-2*s, 1, -5*s),
			grass(-2*s, 1, -6*s),
			grass(-1*s, 1, -6*s),
			grass(0*s, 1, -6*s),
			grass(0*s, 1, -7*s),
			grass(1*s, 1, -7*s),
			grass(2*s, 1, -7*s),
			grass(3*s, 1, -7*s),
			grass(4*s, 1, -7*s),
			grass(5*s, 1, -7*s),
			grass(5*s, 1, -6*s),
			grass(6*s, 1, -6*s),
			grass(6*s, 1, -5*s),
			grass(7*s, 1, -6*s), grass(7*s, 1, -7*s), grass(6*s, 1, -7*s),
			grass(7*s, 1, -5*s),
			grass(7*s, 1, -4*s),
			grass(7*s, 1, -3*s),
			grass(7*s, 1, -2*s),
			grass(6*s, 1, -2*s),
			grass(7*s, 1, -1*s),
			grass(7*s, 1, 0*s),
			grass(7*s, 1, 1*s),
			grass(7*s, 1, 2*s),
			grass(7*s, 1, 3*s), grass(7*s, 1, 4*s), grass(7*s, 1, 5*s), grass(6*s, 1, 4*s), grass(6*s, 1, 5*s), grass(5*s, 1, 5*s), grass(4*s, 1, 5*s),
			grass(6*s, 1, 2*s),
			grass(6*s, 1, 3*s),
			grass(5*s, 1, 3*s),
			grass(5*s, 1, 4*s),
			grass(4*s, 1, 4*s),
			grass(3*s, 1, 4*s),
			grass(3*s, 1, 5*s),
			grass(2*s, 1, 5*s),
			grass(1*s, 1, 5*s),
			grass(0*s, 1, 5*s),
			grass(-1*s, 1, 5*s),
			grass(-2*s, 1, 5*s), grass(-3*s, 1, 5*s), grass(-4*s, 1, 5*s), grass(-3*s, 1, 4*s), grass(-4*s, 1, 4*s), grass(-4*s, 1, 3*s), grass(-4*s, 1, 2*s),
			grass(-1*s, 1, 4*s),
			grass(-2*s, 1, 4*s),
			grass(-2*s, 1, 3*s),
			grass(-3*s, 1, 3*s),
			grass(-3*s, 1, 2*s),
			grass(-3*s, 1, 1*s),
			grass(-4*s, 1, 1*s),
		},
	}
}

func layer1() shapes.Shape {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			grassArea(-4-13, 0, -7, 13),
			grassArea(-4-13, 0, -7-13, 13),
			grassArea(-4, 0, -7-13, 13),
			grassArea(-4+13, 0, -7-13, 13),
			grassArea(-4+13, 0, -7, 13),

			grassArea(-4-13, 0, -7-13*2, 13),
			grassArea(-4, 0, -7-13*2, 13),
			grassArea(-4+13, 0, -7-13*2, 13),

			tree(-5, 1, 3),
			tree(-4, 1, -1),
			tree(-4, 1, -4),
			tree(-1, 1, -7),
			tree(3, 1, -8),
			tree(6, 1, -7),
			tree(8, 1, -5),
			tree(8, 1, -2),

			grass(-4, 0, 0),
			grass(-4, 0, -1),
			grass(-4, 0, -2),
			grass(-4, 0, -3),
			grass(-4, 0, -4),
			grass(-4, 0, -5), grass(-4, 0, -6), grass(-4, 0, -7),
			grass(-3, 0, -4),
			grass(-3, 0, -5),
			grass(-3, 0, -6), grass(-3, 0, -7), grass(-2, 0, -7), grass(-1, 0, -7),
			grass(-2, 0, -5),
			grass(-2, 0, -6),
			grass(-1, 0, -6),
			grass(0, 0, -6),
			grass(0, 0, -7),
			grass(1, 0, -7),
			grass(2, 0, -7),
			grass(3, 0, -7),
			grass(4, 0, -7),
			grass(5, 0, -7),
			grass(5, 0, -6),
			grass(6, 0, -6),
			grass(6, 0, -5),
			grass(7, 0, -6), grass(7, 0, -7), grass(6, 0, -7),
			grass(7, 0, -5),
			grass(7, 0, -4),
			grass(7, 0, -3),
			grass(7, 0, -2),
			grass(6, 0, -2),
			grass(7, 0, -1),
			grass(7, 0, 0),
			grass(7, 0, 1),
			grass(7, 0, 2),
			grass(7, 0, 3), grass(7, 0, 4), grass(7, 0, 5), grass(6, 0, 4), grass(6, 0, 5), grass(5, 0, 5), grass(4, 0, 5),
			grass(6, 0, 2),
			grass(6, 0, 3),
			grass(5, 0, 3),
			grass(5, 0, 4),
			grass(4, 0, 4),
			grass(3, 0, 4),
			grass(3, 0, 5),
			grass(2, 0, 5),
			grass(1, 0, 5),
			grass(0, 0, 5),
			grass(-1, 0, 5),
			grass(-2, 0, 5), grass(-3, 0, 5), grass(-4, 0, 5), grass(-3, 0, 4), grass(-4, 0, 4), grass(-4, 0, 3), grass(-4, 0, 2),
			grass(-1, 0, 4),
			grass(-2, 0, 4),
			grass(-2, 0, 3),
			grass(-3, 0, 3),
			grass(-3, 0, 2),
			grass(-3, 0, 1),
			grass(-4, 0, 1),
		},
	}
}

func layer0() shapes.Shape {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			sand(-3, -1, 0),
			sand(-3, -1, -1),
			sand(-2, -1, -2),
			sand(-1, -1, -3),
			sand(-0, -1, -3),
			sand(-0, -1, -4),
			sand(1, -1, -5),
			sand(2, -1, -6),
			sand(3, -1, -6),
			sand(4, -1, -6),
			sand(5, -1, -5),
			sand(6, -1, -4),
			sand(6, -1, -3),
			sand(5, -1, -2),
			sand(6, -1, -1),
			sand(6, -1, 0),
			sand(6, -1, 1),
			sand(5, -1, 2),
			sand(4, -1, 3),
			sand(3, -1, 3),
			sand(2, -1, 4),
			sand(1, -1, 4),
			sand(0, -1, 4),
			sand(-1, -1, 3),
			sand(-2, -1, 2),
			sand(-2, -1, 1),

			sand(-3, -1, -2),
			sand(-3, -1, -3),
			sand(-3, -1, -4),
			sand(-3, -1, -5),

			sand(-2, -1, -3),
			sand(-2, -1, -4),
			sand(-2, -1, -5),

			sand(-1, -1, -4),
			sand(-1, -1, -5),

			sand(0, -1, -5),
			sand(0, -1, -6),
			sand(1, -1, -6),
		},
	}
}

func torch(x, y, z float64) shapes.Shape {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Cube{
				vec3.Vec3{x, y, z},
				vec3.Vec3{0.1, 0.9, 0.1},
				space.Material_Diffuse{texture.NewTexturemap("res/mc/torch.png", 2.2)},
			},
		},
	}
}

func grassArea(x, y, z float64, size float64) shapes.Shape {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Cube{
				vec3.Vec3{x, y - size - 1, z},
				vec3.Vec3{size, size, size},
				space.Material_Diffuse{texture.NewScaledTexturemap("res/mc/grass.jpg", 2.2, size)},
			},
		},
	}
}

func dirtArea(x, y, z float64, size float64) shapes.Shape {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			shapes.Cube{
				vec3.Vec3{x, y - size - 1, z},
				vec3.Vec3{size, size, size},
				space.Material_Diffuse{texture.NewScaledTexturemap("res/mc/dirt_grass.jpg", 2.2, size)},
			},
		},
	}
}

func tree(x, y, z float64) shapes.Shape {
	return shapes.Group{
		space.NoTransformation(),
		[]shapes.Shape{
			log(x, y, z),
			log(x, y+1, z),
			log(x, y+2, z),
			log(x, y+3, z),
			leave(x, y+4, z), leave(x+1, y+4, z), leave(x-1, y+4, z), leave(x, y+4, z-1), leave(x, y+4, z+1),
			leave(x+1, y+2, z), leave(x+1, y+2, z+1), leave(x, y+2, z+1), leave(x-1, y+2, z+1), leave(x-1, y+2, z), leave(x-1, y+2, z-1), leave(x, y+2, z-1), leave(x+1, y+2, z-1),
			leave(x+2, y+2, z), leave(x+2, y+2, z+1), leave(x+2, y+2, z-1),
			leave(x-2, y+2, z), leave(x-2, y+2, z+1), leave(x-2, y+2, z-1),
			leave(x-1, y+2, z+2), leave(x, y+2, z+2), leave(x+1, y+2, z+2),
			leave(x-1, y+2, z-2), leave(x, y+2, z-2), leave(x+1, y+2, z-2),

			leave(x+1, y+3, z), leave(x+1, y+3, z+1), leave(x, y+3, z+1), leave(x-1, y+3, z+1), leave(x-1, y+3, z), leave(x-1, y+3, z-1), leave(x, y+3, z-1), leave(x+1, y+3, z-1),
			leave(x+2, y+3, z), leave(x+2, y+3, z+1), leave(x+2, y+3, z-1),
			leave(x-2, y+3, z), leave(x-2, y+3, z+1), leave(x-2, y+3, z-1),
			leave(x-1, y+3, z+2), leave(x, y+3, z+2), leave(x+1, y+3, z+2),
			leave(x-1, y+3, z-2), leave(x, y+3, z-2), leave(x+1, y+3, z-2),
		},
	}
}

func leave(x, y, z float64) shapes.Shape {
	return shapes.Cube{
		vec3.Vec3{x, y, z},
		vec3.Vec3{1, 1, 1},
		space.Material_Diffuse{texture.NewTexturemap("res/mc/leaves.png", 2.2)},
	}
}

func log(x, y, z float64) shapes.Shape {
	return shapes.Cube{
		vec3.Vec3{x, y, z},
		vec3.Vec3{1, 1, 1},
		space.Material_Diffuse{texture.NewMCTexturemap("res/mc/log_top.png", "res/mc/log_side.jpg")},
	}
}

func sand(x, y, z float64) shapes.Shape {
	return shapes.Cube{
		vec3.Vec3{x, y, z},
		vec3.Vec3{1, 1, 1},
		space.Material_Diffuse{texture.NewTexturemap("res/mc/sand.jpg", 2.2)},
	}
}

func dirt(x, y, z float64) shapes.Shape {
	return shapes.Cube{
		vec3.Vec3{x, y, z},
		vec3.Vec3{1, 1, 1},
		space.Material_Diffuse{texture.NewTexturemap("res/mc/dirt.jpg", 2.2)},
	}
}

func grass(x, y, z float64) shapes.Shape {
	return shapes.Cube{
		vec3.Vec3{x, y, z},
		vec3.Vec3{1, 1, 1},
		space.Material_Diffuse{texture.NewMCTexturemap("res/mc/dirt.jpg", "res/mc/dirt_grass.jpg", "res/mc/grass.jpg")},
	}
}

func GetMinecraftSceneLight() []lights.Light {
	return []lights.Light{
		lights.Ambient{
			vec3.Vec3{0.1, 0.1, 0.1},
		},
		lights.Point{
			vec3.Vec3{-2.5, 2, 2},
			vec3.Vec3{2, 2, 2},
		},
		lights.Directional{
			vec3.Normalize(vec3.Vec3{-0.1, -1, 0.1}),
			vec3.Vec3{5, 5, 5},
		},
	}
}

func GetMinecraftSceneCamera() camera.Camera {
	return camera.PinholeCamera{
		Position:     vec3.Vec3{5, 2.8, 8},
		Direction:    vec3.Normalize(vec3.Vec3{-0.3, 0, -1}),
		OpeningAngle: math.Pi / 8,
		Width:        200,
		Height:       200,
	}
}

func GetMinecraftSceneCamera2() camera.Camera {
	return camera.PinholeCamera{
		// Position: vec3.Vec3{0, 8, 5},
		Position: vec3.Vec3{0, 40, -10},
		// Direction: vec3.Normalize(vec3.Vec3{0, -0.5, -0.5}),
		Direction:    vec3.Normalize(vec3.Vec3{0, -1, 0}),
		OpeningAngle: math.Pi / 8,
		Width:        200,
		Height:       200,
	}
}
