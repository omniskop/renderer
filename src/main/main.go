package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	"math"
	"os"
	"renderer"
	"renderer/camera"
	"renderer/lights"
	"renderer/shapes"
	"runtime/pprof"
	"space"
	"space/texture"
	"time"
	"vec3"
)

var startTime time.Time

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	startTime = time.Now()

	threads := 8

	fmt.Printf("Using %d threads\n", threads)

	// img := renderer.Render(
	// 	camera.PinholeCamera{
	// 		Position:     vec3.Vec3{-11, 11, 11},
	// 		Direction:    vec3.Normalize(vec3.Vec3{1, -1, -1}),
	// 		OpeningAngle: math.Pi / 4,
	// 		Width:        2000,
	// 		Height:       2000,
	// 	},
	// 	shapes.Group{
	// 		space.NoTransformation(),
	// 		[]shapes.Shape{
	// 			// shapes.NewTriangle(vec3.Vec3{-1, 0, 0}, vec3.Vec3{0, 1, 0}, vec3.Vec3{1, 0, 0}, space.Material_Diffuse{texture.NewColor(1, 0, 0)}),
	// 			shapes.NewTriangleMoller(vec3.Vec3{0, 0, 0}, vec3.Vec3{0, 0, 1}, vec3.Vec3{1, 0, 0}, space.Material_Diffuse{texture.NewColor(1, 0, 0)}),
	// 			// loadShapeFile("cube", "objects/cube.obj"),
	// 			shapes.Sphere{vec3.Vec3{0, 0, 0}, 0.2, space.Material_Diffuse{texture.NewColor(1, 0, 0)}},
	// 			shapes.Sphere{vec3.Vec3{0, 1, 0}, 0.2, space.Material_Diffuse{texture.NewColor(0, 1, 0)}},
	// 			shapes.Background{space.Material_Sky{texture.NewColor(1, 1, 1)}},
	// 		},
	// 	},
	// 	[]lights.Light{},
	// 	100,
	// 	2,
	// 	8,
	// )

	img := renderer.Render(
		camera.PinholeCamera{
			Position:     vec3.Vec3{0, 0, 6},
			Direction:    vec3.Normalize(vec3.Vec3{0, 0, -1}),
			OpeningAngle: math.Pi / 4,
			Width:        100,
			Height:       100,
		},
		shapes.Group{
			space.NoTransformation(),
			[]shapes.Shape{
				shapes.Sphere{vec3.Vec3{-1, 0, 0}, 1, space.Material_Diffuse{texture.NewColor(1, 0, 0)}},
				shapes.Sphere{vec3.Vec3{1, 0, 0}, 1, space.Material_Diffuse{texture.NewColor(0, 1, 0)}},
				shapes.Background{space.Material_Sky{texture.NewColor(1, 1, 1)}},
			},
		},
		[]lights.Light{},
		100,
		5,
		8,
	)

	fmt.Println("Rendering took ", time.Since(startTime))

	file, err := os.Create("renderings/out.png")
	if err != nil {
		log.Fatal("unable to create file")
	}

	err = png.Encode(file, img)

	if err != nil {
		log.Fatal("unable to encode png")
	}
}
