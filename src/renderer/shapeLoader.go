package renderer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"renderer/shapes"
	"space"
	"space/texture"
	"vec3"

	"github.com/udhos/gwob"
)

var parserOptions = gwob.ObjParserOptions{
	LogStats:      false,
	Logger:        func(a string) { fmt.Println(a) },
	IgnoreNormals: false,
}

func loadShapeFile(name string, path string) shapes.Group {
	file, err := os.Open(path)
	if err != nil {
		fmt.Print("An error occoured while reading an obj shapeFile. (" + name + ";" + path + ")")
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	object, err := gwob.NewObjFromReader(name, reader, &parserOptions)
	if err != nil {
		fmt.Print("An error occoured while parsing an obj shapeFile. (" + name + ";" + path + ")")
		log.Fatal(err)
	}

	shapeList := make([]shapes.Shape, len(object.Indices)/3)

	for i, c := 0, 0; i < len(object.Indices); i += 3 {
		shapeList[c] = shapes.NewTriangleMoller(
			vec3.Vec3{
				object.Coord64(object.Indices[i]*3 + 0),
				object.Coord64(object.Indices[i]*3 + 1),
				object.Coord64(object.Indices[i]*3 + 2),
			},
			vec3.Vec3{
				object.Coord64(object.Indices[i+1]*3 + 0),
				object.Coord64(object.Indices[i+1]*3 + 1),
				object.Coord64(object.Indices[i+1]*3 + 2),
			},
			vec3.Vec3{
				object.Coord64(object.Indices[i+2]*3 + 0),
				object.Coord64(object.Indices[i+2]*3 + 1),
				object.Coord64(object.Indices[i+2]*3 + 2),
			},
			space.Material_Diffuse{texture.NewColor(0, 0, 1)},
		)
		c++
	}

	group := shapes.Group{
		Transformation: space.NoTransformation(),
		Shapes:         shapeList,
	}

	return group
}
