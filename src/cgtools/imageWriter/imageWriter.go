package imageWriter

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"math"
)


type imageWriter struct{
    width   int
    height  int
    img     *image.RGBA
}

func New(data []float64, width int, height int) imageWriter {
    img := image.NewRGBA(image.Rect(0, 0, width, height));
    
    // Create a colored image of the given width and height.

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{
				R: uint8(math.Floor(data[(y * width + x) * 3    ] * 255)),
				G: uint8(math.Floor(data[(y * width + x) * 3 + 1] * 255)),
				B: uint8(math.Floor(data[(y * width + x) * 3 + 2] * 255)),
				A: 255,
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









