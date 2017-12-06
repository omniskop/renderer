package scenes

import (
    "customtools/shapes"
    "customtools/vec3"
    "customtools/camera"
    "math"
)


func GetComparisonScene() shapes.Group {
    return shapes.Group{
        []shapes.Shape{
            /*shapes.Sphere{
                vec3.Vec3{-2,0,-30},
                1,
                shapes.Material_Diffuse{vec3.Vec3{0.7,1,0.7}},
            },
            shapes.Sphere{
                vec3.Vec3{0,0,-30},
                1,
                shapes.Material_Metal{vec3.Vec3{1,0.7,0.7}, 0},
            },
            shapes.Sphere{
                vec3.Vec3{2,0,-30},
                1,
                shapes.Material_Diffuse{vec3.Vec3{0.7,0.7,1}},
            },*/
            shapes.Sphere{
                vec3.Vec3{0,-0.5,-15},
                1,
                shapes.Material_Transparent{vec3.Vec3{1,1,1}, 1.3},
            },
            shapes.Sphere{
                vec3.Vec3{-0.9,-0.9,-14},
                0.2,
                shapes.Material_Diffuse{vec3.Vec3{1,0,0}},
            },
            shapes.Plane{vec3.Vec3{0,-1,0}, vec3.Vec3{0,1,0}, shapes.Material_Metal{vec3.Vec3{0.1,0.1,0.4}, 0.05}},
            shapes.Plane{vec3.Vec3{0,-1.1,0}, vec3.Vec3{0,1,0}, shapes.Material_Diffuse{vec3.Vec3{1,0,0}}},
            shapes.Background{shapes.Material_Sky{vec3.Vec3{1,1,1}}},
        },
    }
}

func GetComparisonSceneCamera() camera.Camera {
    return camera.PinholeCamera{
        Position: vec3.Vec3{0,0,0},
        Direction: vec3.Vec3{0,0,-1},
        OpeningAngle: math.Pi / 10,
        Width: 200,
        Height: 200,
    }
}