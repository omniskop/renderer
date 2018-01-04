package scenes

import (
    "math"
    "customtools/shapes"
    "customtools/vec3"
    "customtools/space"
    "customtools/camera"
    "customtools/texture"
)

func GetTextureScene() shapes.Group {
    return shapes.Group{
        space.NoTransformation(),
        []shapes.Shape{
            shapes.Sphere{vec3.Vec3{0,1,0}, 1, space.Material_Texture{ texture.NewTexturemap("res/tag.png",2.2) }},
            // shapes.Plane{vec3.Vec3{0,0,0}, vec3.Vec3{0,1,0}, space.Material_Metal{vec3.Vec3{0.9,0.9,1}, 0.05}},
            shapes.Plane{vec3.Vec3{0,0,0}, vec3.Vec3{0,1,0}, space.Material_Texture{ texture.NewColor(vec3.Vec3{0.9,0.9,1}) }},
            shapes.Background{space.Material_Sky{ texture.NewTexturemap("res/sky.jpg",2.2) }},
        },
    }
}

func GetTextureSceneCamera() camera.Camera {
    return camera.PinholeCamera{
        Position: vec3.Vec3{0,2,10},
        Direction: vec3.Normalize(vec3.Vec3{0,0,-0.5}),
        OpeningAngle: math.Pi / 10,
        Width: 200,
        Height: 200,
    }
}