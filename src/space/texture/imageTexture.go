package texture

import (
	"image"
	_ "image/jpeg" // Add jpeg
	_ "image/png"  // Add png
	"math"
	"os"
	"vec3"
)

type ImageTexture struct {
	image          image.Image
	width          int
	height         int
	pixelBuffer    []uint8
	componentScale float64
	gamma          float64
}

func NewImageTexture(filename string, gamma float64) *ImageTexture {
	out := ImageTexture{}
	infile, err := os.Open(filename)
	if err != nil {
		panic("Error: There is no file with the name '" + filename + "'!")
	}
	defer infile.Close()
	out.image, _, err = image.Decode(infile)
	if err != nil {
		panic("Error: The texture '" + filename + "' could not be decoded as an image.")
	}
	out.width = out.image.Bounds().Dx()
	out.height = out.image.Bounds().Dy()
	out.gamma = gamma

	out.componentScale = 65536

	return &out
}

func (this *ImageTexture) SamplePoint(u, v float64) vec3.Vec3 {
	x := int((u - math.Floor(u)) * float64(this.width))
	y := int((v - math.Floor(v)) * float64(this.height))

	r, g, b, _ := this.image.At(x%this.image.Bounds().Max.X, y%this.image.Bounds().Max.Y).RGBA()
	//reverse gamma
	return vec3.Vec3{
		math.Pow(float64(r)/this.componentScale, this.gamma),
		math.Pow(float64(g)/this.componentScale, this.gamma),
		math.Pow(float64(b)/this.componentScale, this.gamma),
	}
}
