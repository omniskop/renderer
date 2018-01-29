package texture

import (
	"vec3"
)

type mcTexturemap struct {
	image   []*ImageTexture
	scaling float64
}

func NewMCTexturemap(filenames ...string) *mcTexturemap {
	out := mcTexturemap{}
	// out.image = []imageTexture.ImageTexture
	for _, v := range filenames {
		out.image = append(out.image, NewImageTexture(v, 2.2))
	}
	out.scaling = 1
	return &out
}

func (this *mcTexturemap) SamplePoint(x, y, z float64) vec3.Vec3 {
	return this.image[int(z)%len(this.image)].SamplePoint(x*this.scaling, y*this.scaling)
}
