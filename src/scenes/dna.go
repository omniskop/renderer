package scenes

import (
    "customtools/shapes"
    "customtools/vec3"
    "customtools/camera"
    "customtools/space"
    "customtools/texture"
    "math"
)
/*
Position: vec3.Vec3{0,0,4},
Direction: vec3.Normalize(vec3.Vec3{0,-0.01,-1}),
OpeningAngle: math.Pi / 7,
*/
func GetDnaScene() shapes.Group {
    return shapes.Group{
        space.NoTransformation(),
        []shapes.Shape{
            // getDnaSequence(vec3.Vec3{0, 0,0},vec3.Vec3{0,1,0}, rotateAaroundB(vec3.Vec3{5,0,0},vec3.Vec3{0,1,0}, math.Pi * 0)),
            // getDnaSequence(vec3.Vec3{0, 3,0},vec3.Vec3{0,1,0}, rotateAaroundB(vec3.Vec3{5,0,0},vec3.Vec3{0,1,0}, math.Pi * 0.1)),
            // getDnaSequence(vec3.Vec3{0,6,0},vec3.Vec3{0,1,0}, rotateAaroundB(vec3.Vec3{5,0,0},vec3.Vec3{0,1,0}, math.Pi * 0.2)),
            // getDnaSequence(vec3.Vec3{0,9,0},vec3.Vec3{0,1,0}, rotateAaroundB(vec3.Vec3{5,0,0},vec3.Vec3{0,1,0}, math.Pi * 0.3)),
            shapes.Group{space.NoTransformation(),generateDnaString(vec3.Vec3{0,0,-1}, vec3.Vec3{5,0,0},40)},
            shapes.Plane{vec3.Vec3{0,-7,0}, vec3.Vec3{0,1,0}, space.Material_Metal{vec3.Vec3{0.1,0.1,0.4}, 0.05}},
            shapes.Sphere{vec3.Vec3{-15,0,-60}, 10, space.Material_Metal{vec3.Vec3{1,1,1}, 0}},
            shapes.Background{space.Material_Sky{texture.NewColor(vec3.Vec3{1,1,1})}},
        },
    }
}

func generateDnaString(normal, side vec3.Vec3, length int) []shapes.Shape {
    out := make([]shapes.Shape, length)
    var n vec3.Vec3
    for i := 0;i < length; i++ {
        // n = rotateAaroundB(vec3.Add(normal, vec3.Multiply(-0.03,side)),normal, math.Pi * float64(i) * 0.1)
        n = vec3.Add(normal, vec3.Vec3{float64(i) * 0.01, float64(i) * 0.005, 0})
        out[i] = getDnaSequence(vec3.Multiply(float64(i * 3), n),n, rotateAaroundB(side,n, math.Pi * float64(i) * 0.1))
    }
    return out
}

func getDnaSequence(position vec3.Vec3, normal vec3.Vec3, side vec3.Vec3) shapes.Group {
    return shapes.Group{
        space.NoTransformation(),
        []shapes.Shape{
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.09), vec3.Multiply(4,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.1), vec3.Multiply(3,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.7,side)),normal, math.Pi * 0.05), vec3.Multiply(3,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{//large
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.05), vec3.Multiply(3,normal)),
                1.5,space.Material_Diffuse{vec3.Vec3{0.7,0,.5}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.3,side)),normal, math.Pi * 0), vec3.Multiply(2,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.4,side)),normal, math.Pi * 0), vec3.Multiply(1.2,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.2,side)),normal, math.Pi * 0.02), vec3.Multiply(1,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{ // caps
                Position: vec3.Subtract(position, side),
                Material: space.Material_Diffuse{vec3.Red},
                Radius: 1,
            },
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.9,side)),normal, math.Pi * 0.02), vec3.Multiply(-0.1,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.8,side)),normal, math.Pi * 1.98), vec3.Multiply(0,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.7,side)),normal, math.Pi * 0.1), vec3.Multiply(0.2,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.5,side)),normal, math.Pi * 0), vec3.Multiply(0,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.3,side)),normal, math.Pi * -0.3), vec3.Multiply(0,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.3,side)),normal, math.Pi * 0.1), vec3.Multiply(0.3,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{ // center
                position,
                1,
                space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}},
            },
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.3,side)),normal, math.Pi * 0.1), vec3.Multiply(0.2,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.3,side)),normal, math.Pi * -0.1), vec3.Multiply(-0.3,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.5,side)),normal, math.Pi * 0), vec3.Multiply(0,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.7,side)),normal, math.Pi * 0.1), vec3.Multiply(0,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.8,side)),normal, math.Pi * 1.93), vec3.Multiply(0,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.9,side)),normal, math.Pi * 0.02), vec3.Multiply(0.2,normal)),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{ // caps
                Position:  vec3.Add(position, side),
                Material: space.Material_Diffuse{vec3.Red},
                Radius: 1,
            },
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.2,side)),normal, math.Pi * 0.02), vec3.Multiply(1,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.4,side)),normal, math.Pi * 0), vec3.Multiply(1.2,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.3,side)),normal, math.Pi * 0), vec3.Multiply(2,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{//large
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.05), vec3.Multiply(3,normal)),
                1.5,space.Material_Diffuse{vec3.Vec3{0.7,0,.5}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.7,side)),normal, math.Pi * 0.1), vec3.Multiply(3,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.1), vec3.Multiply(3,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.09), vec3.Multiply(4,normal)),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
        },
    }
}

func GetDnaSceneCamera() camera.Camera {
    // return camera.PinholeCamera{
    //     Position: vec3.Vec3{0,100,100},
    //     Direction: vec3.Normalize(vec3.Vec3{0,-1,-0.5}),
    //     OpeningAngle: math.Pi / 15,
    //     Width: 200,
    //     Height: 200,
    // }
    return camera.PinholeCamera{
        Position: vec3.Vec3{0,100,100},
        Direction: vec3.Normalize(vec3.Vec3{0,-0.5,-0.5}),
        OpeningAngle: math.Pi / 10,
        Width: 200,
        Height: 200,
    }
}


func rotateAaroundB(v, k vec3.Vec3, angle float64) vec3.Vec3 {
    return vec3.Add(
        vec3.Multiply(math.Cos(angle),v),
        vec3.Multiply(math.Sin(angle), vec3.CrossProduct(k,v)),
        vec3.Multiply(vec3.DotProduct(k,v) * (1-math.Cos(angle)),k),
    )
}