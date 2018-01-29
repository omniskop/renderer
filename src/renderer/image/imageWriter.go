package image

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type imageWriter struct {
	width  int
	height int
	img    *image.RGBA64
}

func NewImageWriter(data []float64, width int, height int) imageWriter {
	img := image.NewRGBA64(image.Rect(0, 0, width, height))

	// Create a colored image of the given width and height.

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA64{
				R: uint16(math.Floor(data[(y*width+x)*3] * 65535)),
				G: uint16(math.Floor(data[(y*width+x)*3+1] * 65535)),
				B: uint16(math.Floor(data[(y*width+x)*3+2] * 65535)),
				A: 65535,
			})
		}
	}

	return imageWriter{
		width,
		height,
		img,
	}
}

func (iw imageWriter) Write(filename string) error {
	f, err := os.Create(filename)
	defer f.Close()

	if err != nil {
		return err
	}

	if err := png.Encode(f, iw.img); err != nil {
		// f.Close()
		// log.Fatal(err)
		return err
	}

	// if err := f.Close(); err != nil {
	//     log.Fatal(err)
	// }
	return nil
}
