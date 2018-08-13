package renderer

import (
	"fmt"
	"image"
	"math"
	"math/rand"
	"renderer/camera"
	"time"

	"image/color"
	"renderer/lights"
	"renderer/shapes"
	"runtime"
	"space"
	"space/ray"
	"sync"
	"vec3"
)

type renderMeta struct {
	progress float32
}

func newRenderMeta() *renderMeta {
	return &renderMeta{
		progress: 0,
	}
}

// Render an image
func Render(cam camera.Camera, scene shapes.Group, light []lights.Light, subsamples int, depth int, threads int) image.Image {
	runtime.GOMAXPROCS(threads)

	samplingPointsPerThread := subsamples / threads

	images := make([]image.Image, threads)
	renderMetas := make([]*renderMeta, threads)

	var wg sync.WaitGroup
	channel := make(chan image.Image)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		renderMetas[i] = newRenderMeta()
		go renderThread(&wg, channel, cam, scene, light, samplingPointsPerThread, depth, renderMetas[i])
	}

	progressStop := make(chan bool)
	go progressPrinter(renderMetas, progressStop)

	for i := 0; i < threads; i++ {
		images[i] = <-channel
	}

	wg.Wait()

	progressStop <- true

	finalImage := image.NewRGBA(image.Rect(0, 0, cam.GetWidth(), cam.GetHeight()))

	threadsFloat := float64(threads)
	for x := 0; x < cam.GetWidth(); x++ {
		for y := 0; y < cam.GetHeight(); y++ {
			c := vec3.Black
			for _, img := range images {
				r, g, b, _ := img.At(x, y).RGBA()
				c.Add(vec3.Vec3{float64(r) / 65535, float64(g) / 65535, float64(b) / 65535})
			}
			c = processColor(vec3.Divide(c, threadsFloat), 2.2)
			finalImage.Set(x, y, color.RGBA{uint8(c.X * 255), uint8(c.Y * 255), uint8(c.Z * 255), 255})
		}
	}

	return finalImage
}

func progressPrinter(metas []*renderMeta, stop chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	// fmt.Print("\n")
	for {
		select {
		case <-stop:
			ticker.Stop()
			// fmt.Print("\r\n")
			return
		case <-ticker.C:
			var progress float32
			for _, meta := range metas {
				// i know this is a data race
				progress += meta.progress
			}

			progress /= float32(len(metas))
			fmt.Printf("\rRendering... %d%%", int(progress*100))
		}
	}
}

func deg2rad(x float64) float64 {
	return (x / 360) * math.Pi * 2
}

func renderThread(wg *sync.WaitGroup, out chan image.Image, cam camera.Camera, scene shapes.Shape, light []lights.Light, sPoints, depth int, meta *renderMeta) {
	img := raytrace(cam, scene, light, sPoints, depth, meta)
	out <- img
	wg.Done()
}

func raytrace(cam camera.Camera, shapes shapes.Shape, light []lights.Light, sPoints, depth int, meta *renderMeta) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, cam.GetWidth(), cam.GetHeight()))
	pointPerPixelAxis := int(math.Sqrt(float64(sPoints)))
	for x := 0; x < cam.GetWidth(); x++ {
		for y := 0; y < cam.GetHeight(); y++ {
			img.Set(
				x,
				y,
				getColorForPixel(shapes, cam, light, x, y, pointPerPixelAxis, depth).Color(),
			)
		}
		percent := float32(x) / float32(cam.GetWidth())
		meta.progress = percent
	}

	return img
}

func getColorForPixel(shapes shapes.Shape, cam camera.Camera, light []lights.Light, pX, pY, supersamplingPoints, depth int) vec3.Vec3 {
	color := vec3.Vec3{0, 0, 0}
	for x := 0; x < supersamplingPoints; x++ {
		for y := 0; y < supersamplingPoints; y++ {
			pixelRay := cam.GetRayForPixel(
				float64(pX)+(float64(x)+rand.Float64())/float64(supersamplingPoints),
				float64(pY)+(float64(y)+rand.Float64())/float64(supersamplingPoints),
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
		fmt.Println("Sollte nicht passieren.")
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
			fmt.Println("Sollte nicht passieren.")
			break
		}

		emission := closestHit.Material.EmittedRadiance(*currentRay, *closestHit)
		currentRay = closestHit.Material.ScatteredRay(*currentRay, *closestHit)

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
