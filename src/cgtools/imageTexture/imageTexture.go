package imageTexture

import (
	"customtools/vec3"
	"image"
	_ "image/jpeg" // Add jpeg
	_ "image/png"  // Add png
	"math"
	"os"
)

type ImageTexture struct {
	image          image.Image
	width          int
	height         int
	pixelBuffer    []uint8
	componentScale float64
	gamma          float64
}

func New(filename string, gamma float64) *ImageTexture {
	out := ImageTexture{}
	infile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer infile.Close()
	out.image, _, err = image.Decode(infile)
	if err != nil {
		panic(err)
	}
	out.width = out.image.Bounds().Dx()
	out.height = out.image.Bounds().Dy()
	out.gamma = gamma
	// img, ok := out.image.(*image.RGBA)
	// if !ok {
	//     panic("image > imageBuffer")
	// }
	// out.pixelBuffer = img.Pix
	out.componentScale = 65536

	return &out
}

func (this *ImageTexture) SamplePoint(u, v float64) vec3.Vec3 {
	// int x = (int) ((u - Math.floor(u)) * width);
	x := int((u - math.Floor(u)) * float64(this.width))
	// int y = (int) ((v - Math.floor(v)) * height);
	y := int((v - math.Floor(v)) * float64(this.height))
	// image.getRaster().getPixel(x, y, pixelBuffer);

	// Vec3 color = vec3(pixelBuffer[0], pixelBuffer[1], pixelBuffer[2]);
	// return divide(color, componentScale);

	// return vec3.Vec3{
	//     float64(this.pixelBuffer[(y * this.width + x)*4 + 0]) * this.componentScale,
	//     float64(this.pixelBuffer[(y * this.width + x)*4 + 1]) * this.componentScale,
	//     float64(this.pixelBuffer[(y * this.width + x)*4 + 2]) * this.componentScale,
	// }
	r, g, b, _ := this.image.At(x%this.image.Bounds().Max.X, y%this.image.Bounds().Max.Y).RGBA()
	//reverse gamma
	return vec3.Vec3{
		math.Pow(float64(r)/this.componentScale, this.gamma),
		math.Pow(float64(g)/this.componentScale, this.gamma),
		math.Pow(float64(b)/this.componentScale, this.gamma),
	}
}

/*
public class ImageTexture {
    private BufferedImage image;
    public final int width;
    public final int height;
    private final double[] pixelBuffer;
    private final double componentScale;

    public ImageTexture(String filename) throws IOException {
        image = ImageIO.read(new File(filename));
        width = image.getWidth();
        height = image.getHeight();
        pixelBuffer = new double[image.getRaster().getNumBands()];

        switch (image.getSampleModel().getDataType()) {
        case DataBuffer.TYPE_BYTE:
            componentScale = 255;
            break;
        case DataBuffer.TYPE_USHORT:
            componentScale = 65535;
            break;
        default:
            componentScale = 1;
            break;
        }
    }

    public Vec3 samplePoint(double u, double v) {
        int x = (int) ((u - Math.floor(u)) * width);
        int y = (int) ((v - Math.floor(v)) * height);
        image.getRaster().getPixel(x, y, pixelBuffer);
        Vec3 color = vec3(pixelBuffer[0], pixelBuffer[1], pixelBuffer[2]);
        return divide(color, componentScale);
    }
}*/
