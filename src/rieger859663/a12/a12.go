// package rieger859663
package main

import (
	"cgtools/image"
	"cgtools/random"
	"customtools/camera"
	"customtools/lights"
	"customtools/ray"
	"customtools/shapes"
	"customtools/space"
	"customtools/texture"
	"customtools/vec3"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"runtime/pprof"
	"scenes"
	"sync"
	"time"
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
	random.InitSeed()

	// sceneCamera := scenes.GetDnaSceneCamera()
	sceneCamera := scenes.GetMinecraftSceneCamera()
	// sceneCamera := scenes.GetHumanDnaCamera()

	_ = scenes.GetWaterMolecule()
	_ = space.NewTransformation
	_ = sceneCamera
	_ = texture.NewColor

	//Standart Scenen
	img := multithreadMagic(
		camera.PinholeCamera{
			Position:     sceneCamera.GetPosition(),
			Direction:    sceneCamera.GetDirection(),
			OpeningAngle: math.Pi / 2,
			Width:        1280,
			Height:       720,
		},
		scenes.GetMinecraftScene(),
		scenes.GetMinecraftSceneLight(),
		100,
		10,
	)

	// Kamera tests
	// img := multithreadMagic(
	//     camera.PinholeCamera{
	//         Position: vec3.Vec3{2,1,-10},
	//         Direction: vec3.Normalize(vec3.Vec3{-0.4,-0.3,-1}),
	//         Tilt: 0,
	//         OpeningAngle: math.Pi / 7,
	//         Width: 200,
	//         Height: 200,
	//     },
	//     scenes.GetComparisonScene(),
	//     100,
	//     30,
	// )

	// Performance tests
	// img := raytrace(
	//     camera.PinholeCamera{
	//         Position: sceneCamera.GetPosition(),
	//         Direction: sceneCamera.GetDirection(),
	//         OpeningAngle: sceneCamera.GetOpeningAngle(),
	//         Width: 200,
	//         Height: 200,
	//     },
	//     scenes.GetCylinderScene(),
	//     300,
	//     30,
	// )

	// scene tests
	// img := multithreadMagic(
	//     camera.PinholeCamera{
	//             Position: vec3.Vec3{0,5,5},
	//             Direction: vec3.Normalize(vec3.Vec3{0,-1,-1}),
	//             OpeningAngle: math.Pi / 2,
	//             Width: 400,
	//             Height: 400,
	//         },
	//         shapes.Group{
	//             space.NoTransformation(),
	//             []shapes.Shape{
	// 				shapes.Union(
	// 					shapes.Sphere{
	// 						vec3.Vec3{-0.5,0,0},
	// 						1,
	// 						space.Material_Diffuse{texture.NewColor(1,0,0)},
	// 					},
	// 					shapes.Sphere{
	// 						vec3.Vec3{0.5,0,0},
	// 						1,
	// 						space.Material_Diffuse{texture.NewColor(0,0,1)},
	// 					},
	// 				),
	// 				shapes.Background{space.Material_Sky{texture.NewColor(1,1,1)}},
	// 			},
	//         },
	//         70,
	//         2,
	// )

	log.Print("Rendering took ", time.Since(startTime))

	err := img.Write("doc/cgg-competition-ws-17-859663.png")
	if err != nil {
		log.Print("An error occoured while writing the file.")
		log.Fatal(err)
	} else {
		log.Print("File saved.")
	}
}

func deg2rad(x float64) float64 {
	return (x / 360) * math.Pi * 2
}

func multithreadMagic(cam camera.Camera, scene shapes.Shape, light []lights.Light, supersamplingPoints, depth int) image.Image {
	// threads := runtime.NumCPU()
	threads := 7
	runtime.GOMAXPROCS(threads)
	log.Print(threads, " threads")

	samplingPointsPerThread := supersamplingPoints / threads
	log.Print(samplingPointsPerThread*threads, " sampling points")
	log.Print(samplingPointsPerThread, " sampling points per thread")

	images := make([]image.Image, threads)

	var wg sync.WaitGroup
	channel := make(chan image.Image)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go renderThread(&wg, channel, cam, scene, light, samplingPointsPerThread, depth)
	}

	for i := 0; i < threads; i++ {
		images[i] = <-channel
	}

	wg.Wait()

	finalImage := image.New(cam.GetWidth(), cam.GetHeight())

	threadsFloat := float64(threads)
	for x := 0; x < cam.GetWidth(); x++ {
		for y := 0; y < cam.GetHeight(); y++ {
			color := vec3.Black
			for _, img := range images {
				color.Add(img.GetPixel(x, y))
			}
			finalImage.SetPixel(x, y, processColor(vec3.Divide(color, threadsFloat), 2.2))
		}
	}

	return finalImage
}

func renderThread(wg *sync.WaitGroup, out chan image.Image, cam camera.Camera, scene shapes.Shape, light []lights.Light, sPoints, depth int) {
	img := raytrace(cam, scene, light, sPoints, depth)
	out <- img
	wg.Done()
}

func raytrace(cam camera.Camera, shapes shapes.Shape, light []lights.Light, sPoints, depth int) image.Image {
	img := image.New(cam.GetWidth(), cam.GetHeight())
	pointPerPixelAxis := int(math.Sqrt(float64(sPoints)))
	for x := 0; x < cam.GetWidth(); x++ {
		for y := 0; y < cam.GetHeight(); y++ {
			img.SetPixel(
				x,
				y,
				// processColor(
				getColorForPixel(shapes, cam, light, x, y, pointPerPixelAxis, depth),
				// 2.2,
				// ),

			)
		}
		percent := (float64(x) / float64(cam.GetWidth())) * 100
		if percent == math.Trunc(percent) {
			fmt.Printf("\rRendering... %d%%", int(percent))
		}
	}

	fmt.Printf("\n")

	return img
}

func getColorForPixel(shapes shapes.Shape, cam camera.Camera, light []lights.Light, pX, pY, supersamplingPoints, depth int) vec3.Vec3 {
	color := vec3.Vec3{0, 0, 0}
	for x := 0; x < supersamplingPoints; x++ {
		for y := 0; y < supersamplingPoints; y++ {
			pixelRay := cam.GetRayForPixel(
				float64(pX)+(float64(x)+random.Float64())/float64(supersamplingPoints),
				float64(pY)+(float64(y)+random.Float64())/float64(supersamplingPoints),
			)

			c := calculateRadiance(shapes, light, pixelRay, depth)
			// c := calculateRadiance2(shapes, light, pixelRay, depth)
			color.Add(c)
		}
	}

	color.Divide(float64(supersamplingPoints * supersamplingPoints))

	return color
}

func calculateRadiance(scene shapes.Shape, light []lights.Light, renderRay ray.Ray, depth int) vec3.Vec3 {
	closestHit := scene.Intersect(renderRay)

	radiance := vec3.Zero

	if closestHit == nil {
		log.Print("Sollte nicht passieren.")
		return radiance
	}

	emission := closestHit.Material.EmittedRadiance(renderRay, *closestHit)
	scattered := closestHit.Material.ScatteredRay(renderRay, *closestHit)

	// emission = vec3.Add(emission, lights.SampleLights(scene, light, closestHit))

	if scattered != nil {
		depth--
		if depth != 0 {
			directLight := vec3.Vec3{0, 0, 0}
			if _, ok := closestHit.Material.(space.Material_Diffuse); ok {
				directLight = lights.SampleLights(scene, light, closestHit)
			}
			radiance = vec3.Clamp(vec3.Add(
				directLight,
				calculateRadiance(scene, light, *scattered, depth),
			))
		}
		return vec3.Add(
			emission,
			vec3.MultiplyByVec3(
				closestHit.Material.Albedo(renderRay, *closestHit),
				radiance,
			),
		)
	}
	return emission
}

func calculateRadiance2(scene shapes.Shape, light []lights.Light, renderRay ray.Ray, depth int) vec3.Vec3 {
	currentRay := &renderRay
	value := vec3.Zero
	add := vec3.One
	for i := 0; i < depth; i++ {
		closestHit := scene.Intersect(*currentRay)
		if closestHit == nil {
			log.Print("Sollte nicht passieren.")
			break
		}

		// a, e := closestHit.Material.(space.Material_Diffuse)
		// if e && a.Color.Equals(vec3.Vec3{1,0,0}) && closestHit.Normal.Y < -000.1 {
		// if closestHit.Normal.Y < 0 && true {
		// log.Print("Affensalat! ", closestHit.Normal)

		// value.Add(vec3.Vec3{0,0,1})
		// return value

		// runtime.Breakpoint()
		// value.Add(vec3.MultiplyByVec3(add, vec3.Vec3{.5,0,1}))
		// break;
		// }

		emission := closestHit.Material.EmittedRadiance(*currentRay, *closestHit)
		// c := closestHit.Position.Y / 1
		// if c < 0 {
		//     c = 0
		// }
		// emission := vec3.Vec3{c,c,c}
		currentRay = closestHit.Material.ScatteredRay(*currentRay, *closestHit)

		// currentRay = nil
		// v := closestHit.T * 10
		// // v = 10 - (v/10)*10
		// emission = vec3.Vec3{v, v, v}

		emission = vec3.Add(emission, lights.SampleLights(scene, light, closestHit))

		value.Add(vec3.MultiplyByVec3(add, emission))
		value.Clamp()
		if currentRay != nil {
			add.MultiplyByVec3(closestHit.Material.Albedo(renderRay, *closestHit))
		} else {
			break
		}
	}
	return value
}

func processColor(color vec3.Vec3, gamma float64) vec3.Vec3 {
	return vec3.Vec3{
		math.Pow(color.X, 1/gamma),
		math.Pow(color.Y, 1/gamma),
		math.Pow(color.Z, 1/gamma),
	}
}
