package rieger859663;

import cgtools.Vec3;
import cgtools.ImageWriter;
import java.io.IOException;

public class Image {
    int width;
    int height;
    double imageData[];
    
    public Image(int width, int height) {
        this.width = width;
        this.height = height;
        imageData = new double[width * height * 3];
    }

    public void setPixel(int x, int y, Vec3 color) {
        imageData[y * 3 * width + x * 3    ] = color.x;
        imageData[y * 3 * width + x * 3 + 1] = color.y;
        imageData[y * 3 * width + x * 3 + 2] = color.z;
    }

    public void write(String filename) throws IOException {
        ImageWriter writer = new ImageWriter( imageData, width, height );
        writer.write( filename );
    }
}
