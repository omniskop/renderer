package scenes

import (
    "customtools/shapes"
    "customtools/vec3"
    "customtools/camera"
    "customtools/space"
    "math"
)

func GetCylinderScene() shapes.Group {
    return shapes.Group {
        space.NoTransformation(),
        []shapes.Shape{
            shapes.NewCylinder(
                vec3.Vec3{0,-1.3,-10},
                vec3.Normalize(vec3.Vec3{0,1,0}),
                1,
                2.5,
                space.Material_Metal{vec3.Vec3{1,0.8,0.8}, 0},
                // space.Material_Transparent{vec3.White, 1.5},
            ),
            shapes.NewCylinder(
                vec3.Vec3{-1,-1.3,-13},
                vec3.Normalize(vec3.Vec3{-1,1,-1}),
                0.5,
                2,
                space.Material_Diffuse{vec3.Red},
            ),
            shapes.NewCylinder(
                vec3.Vec3{1,-1.3,-13},
                vec3.Normalize(vec3.Vec3{1,1,-1}),
                0.5,
                2,
                space.Material_Diffuse{vec3.Red},
            ),
            shapes.Sphere{
                vec3.Vec3{0,-1.3,-1.5},
                0.5,
                space.Material_Diffuse{vec3.Vec3{0,1,0}},
            },
            shapes.Sphere{
                vec3.Vec3{0,-1.3,-6},
                0.5,
                space.Material_Diffuse{vec3.Vec3{0,0,1}},
            },
            shapes.Sphere{
                vec3.Vec3{0,0.6,-0.5},
                0.8,
                space.Material_Transparent{vec3.Vec3{1,1,1}, 1.5},
            },
            shapes.Plane{ // bottom
                Position: vec3.Vec3{0,-1.3,0},
                Normal: vec3.Vec3{0,1,0},
                Material: space.Material_Diffuse_Checkerboard{1,vec3.White, vec3.Black},
            },
            shapes.Background{space.Material_Sky{vec3.Vec3{0.8,0.8,1}}},
        },
    }
}

func GetCylinderSceneCamera() camera.PinholeCamera {
    return camera.PinholeCamera{
        Position: vec3.Vec3{0,2,7},
        Direction: vec3.Normalize(vec3.Vec3{0,-0.2,-1}),
        OpeningAngle: math.Pi / 9,
        Width: 500,
        Height: 500,
    }
}