package rieger859663.a01;

import cgtools.Vec3;
import static cgtools.Vec3.*;
import java.io.IOException;
import rieger859663.Image;

public class Main {
    static int width = 160;
    static int height = 90;

    public static void main(String[] args) {
        Image image = new Image(width, height);

        for (int x = 0; x != width; x++) {
            for (int y = 0; y != height; y++) {
                image.setPixel(x, y, pixelColor(x, y));
            }
        }

        String filename = "../doc/a01-checkerboard.png";
        try {
            image.write(filename);
            System.out.println("Wrote image: " + filename);
        } catch (IOException error) {
            System.out.println(String.format("Something went wrong writing: %s: %s", filename, error));
        }
    }

    static Vec3 pixelColor(int x, int y) {
        if(((x / 10) % 2 == 0) ^ ((y / 10) % 2 == 0)) {
            return vec3(0,0,1);
        } else {
            return vec3((double) x / width,1 - (double) x / width, 0);
        }
    }
}
