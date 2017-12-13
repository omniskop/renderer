package scenes

import (
    "customtools/shapes"
    "customtools/vec3"
    "customtools/camera"
    "customtools/space"
    "math"
)


func GetHumanDnaScene() shapes.Group {
    var scale float64 = 11
    return shapes.Group{
        space.NoTransformation(),
        []shapes.Shape{
            
            
            //Linkes Bein
            shapes.NewCylinder(vec3.Multiply(scale,vec3.Vec3{-1,10,0}), vec3.Normalize(vec3.Vec3{-.3,-1,0}),scale * 1, scale * 10, space.Material_Diffuse{vec3.Rgb(38, 50, 175)}),
            //Rechtes Bein
            shapes.NewCylinder(vec3.Multiply(scale,vec3.Vec3{1,10,0}), vec3.Normalize(vec3.Vec3{.3,-1,0}),scale * 1, scale * 10, space.Material_Diffuse{vec3.Rgb(38, 50, 175)}),
            //Körper
            shapes.NewCylinder(vec3.Multiply(scale,vec3.Vec3{0,10,0}), vec3.Normalize(vec3.Vec3{0,1,0}),scale * 3, scale * 10, space.Material_Diffuse{vec3.Rgb(116, 166, 242)}),
            //Hals
            shapes.NewCylinder(vec3.Multiply(scale,vec3.Vec3{0,20,0}), vec3.Normalize(vec3.Vec3{0,1,0}),scale * 1, scale * 1, space.Material_Diffuse{vec3.Rgb(255, 225, 145)}),
            //Kopf
            shapes.Sphere{vec3.Multiply(scale,vec3.Vec3{0,21 + 2,0}), scale * 2, space.Material_Diffuse{vec3.Rgb(255, 225, 145)}},
            //Linkes Auge
            shapes.Sphere{vec3.Multiply(scale,vec3.Vec3{-1,21 + 2,2}), scale * 0.3, space.Material_Diffuse{vec3.Blue}},
            //Rechtes Auge
            shapes.Sphere{vec3.Multiply(scale,vec3.Vec3{1,21 + 2,2}), scale * 0.3, space.Material_Diffuse{vec3.Blue}},
            //Hut
            shapes.NewCone(vec3.Multiply(scale,vec3.Vec3{0,24,0}), vec3.Vec3{0,1,0}, scale * 1.8, scale * 5, space.Material_Diffuse{vec3.Green}),
            
            //Linker Arm
            shapes.NewCylinder(vec3.Multiply(scale,vec3.Vec3{-3,19,0}), vec3.Normalize(vec3.Vec3{-.7,-1,0}),scale * 1, scale * 10, space.Material_Diffuse{vec3.Rgb(255, 225, 145)}),
            //Rechter Arm
            shapes.NewCylinder(vec3.Multiply(scale,vec3.Vec3{3,19,0}), vec3.Normalize(vec3.Vec3{1,-.2,0}),scale * 1, scale * 10, space.Material_Diffuse{vec3.Rgb(255, 225, 145)}),
            
            //Lupe
            shapes.Group{
                space.NewTransformation(vec3.Vec3{-50,-420,0}, 0,0,0),
                // space.NoTransformation(),
                []shapes.Shape{
                shapes.Group{
                    space.NewTransformation(vec3.Vec3{0,0,0}, math.Pi, 0, math.Pi * -0.1),
                    // space.NoTransformation(),
                    []shapes.Shape{
                        shapes.Group{space.NewTransformation(vec3.Vec3{0,-20,0},0,0,0),generateHumanDnaString(vec3.Vec3{50,130,0},vec3.Vec3{1,0,0}, vec3.Vec3{0,0,-5},20)},
                        shapes.NewCone(vec3.Multiply(scale,vec3.Vec3{8,19 - 4,0}), vec3.Vec3{0,1,0}, scale * 3, scale * 2, space.Material_Diffuse{vec3.Red}),
                        // shapes.Sphere{vec3.Vec3{8 * scale,130,0}, scale * 4, space.Material_Transparent{vec3.Vec3{0.9,0.9,1}, 1.5}},
                    },
                },
            },
            },
            
            // shapes.Sphere{vec3.Multiply(scale,vec3.Vec3{0,3,7}), scale * 3, space.Material_Transparent{vec3.White, 1.5}},
            shapes.Plane{vec3.Vec3{0,0,0}, vec3.Normalize(vec3.Vec3{0,1,0}), space.Material_Metal{vec3.Vec3{0.9,0.9,0.9},0.1}},
            // shapes.Sphere{vec3.Vec3{0,0,0}, 1, space.Material_Diffuse{vec3.Vec3{0,1,0}}},
            shapes.Background{space.Material_Sky{vec3.Vec3{1,1,1}}},
        },
    }
}

func generateHumanDnaString(shift,normal, side vec3.Vec3, length int) []shapes.Shape {
    out := make([]shapes.Shape, length)
    var n vec3.Vec3
    for i := 0;i < length; i++ {
        // n = rotateAaroundB(vec3.Add(normal, vec3.Multiply(-0.03,side)),normal, math.Pi * float64(i) * 0.1)
        n = vec3.Add(normal, vec3.Vec3{float64(i) * 0.01, float64(i) * 0.005, 0})
        out[i] = getHumanDnaSequence(shift,vec3.Multiply(float64(i * 3), n),n, rotateAaroundB(side,n, math.Pi * float64(i) * 0.1))
    }
    return out
}

func getHumanDnaSequence(shift, position vec3.Vec3, normal vec3.Vec3, side vec3.Vec3) shapes.Group {
    return shapes.Group{
        space.NoTransformation(),
        []shapes.Shape{
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.09), vec3.Multiply(4,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.1), vec3.Multiply(3,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.7,side)),normal, math.Pi * 0.05), vec3.Multiply(3,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{//large
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.05), vec3.Multiply(3,normal),shift),
                1.5,space.Material_Diffuse{vec3.Vec3{0.7,0,.5}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.3,side)),normal, math.Pi * 0), vec3.Multiply(2,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.4,side)),normal, math.Pi * 0), vec3.Multiply(1.2,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(1.2,side)),normal, math.Pi * 0.02), vec3.Multiply(1,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{ // caps
                Position: vec3.Add(vec3.Subtract(position, side),shift),
                Material: space.Material_Diffuse{vec3.Red},
                Radius: 1,
            },
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.9,side)),normal, math.Pi * 0.02), vec3.Multiply(-0.1,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.8,side)),normal, math.Pi * 1.98), vec3.Multiply(0,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.7,side)),normal, math.Pi * 0.1), vec3.Multiply(0.2,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.5,side)),normal, math.Pi * 0), vec3.Multiply(0,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.3,side)),normal, math.Pi * -0.3), vec3.Multiply(0,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Subtract(position, vec3.Multiply(0.3,side)),normal, math.Pi * 0.1), vec3.Multiply(0.3,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{ // center
                vec3.Add(position,shift),
                1,
                space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}},
            },
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.3,side)),normal, math.Pi * 0.1), vec3.Multiply(0.2,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.3,side)),normal, math.Pi * -0.1), vec3.Multiply(-0.3,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.5,side)),normal, math.Pi * 0), vec3.Multiply(0,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.7,side)),normal, math.Pi * 0.1), vec3.Multiply(0,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.8,side)),normal, math.Pi * 1.93), vec3.Multiply(0,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(0.9,side)),normal, math.Pi * 0.02), vec3.Multiply(0.2,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{1,0.2,0.2}}},
            shapes.Sphere{ // caps
                Position:  vec3.Add(position, side,shift),
                Material: space.Material_Diffuse{vec3.Red},
                Radius: 1,
            },
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.2,side)),normal, math.Pi * 0.02), vec3.Multiply(1,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.4,side)),normal, math.Pi * 0), vec3.Multiply(1.2,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.3,side)),normal, math.Pi * 0), vec3.Multiply(2,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{//large
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.05), vec3.Multiply(3,normal),shift),
                1.5,space.Material_Diffuse{vec3.Vec3{0.7,0,.5}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.7,side)),normal, math.Pi * 0.1), vec3.Multiply(3,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.1), vec3.Multiply(3,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
            shapes.Sphere{
                vec3.Add(rotateAaroundB(vec3.Add(position, vec3.Multiply(1.5,side)),normal, math.Pi * 0.09), vec3.Multiply(4,normal),shift),
                1,space.Material_Diffuse{vec3.Vec3{0.7,0,0}}},
        },
    }
}

func GetHumanDnaCamera() camera.Camera {
    return camera.PinholeCamera{
        Position: vec3.Vec3{0,100,1000},
        Direction: vec3.Normalize(vec3.Vec3{0,0.05,-1}),
        OpeningAngle: math.Pi / 9,
        Width: 200,
        Height: 200,
    }
}

func GetHumanDnaCamera2() camera.Camera {
    return camera.PinholeCamera{
        Position: vec3.Vec3{700,700,800},
        Direction: vec3.Normalize(vec3.Vec3{-1,-0.7,-.9}),
        OpeningAngle: math.Pi / 11,
        Width: 200,
        Height: 200,
    }
}