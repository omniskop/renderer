package space

import (
	"cgtools/random"
	"customtools/ray"
	"customtools/texture"
	"customtools/vec3"
	"math"
)

func RandomDirection() vec3.Vec3 {
	direction := vec3.One
	for direction.Length() > 1 {
		direction = vec3.Vec3{
			random.Float64()*2 - 1,
			random.Float64()*2 - 1,
			random.Float64()*2 - 1,
		}
	}
	return direction
}

type Material interface {
	EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3
	ScatteredRay(r ray.Ray, h Hit) *ray.Ray
	Albedo(r ray.Ray, h Hit) vec3.Vec3
	String() string
}

/*  ======  DIFFUSE  ====== */

type Material_Diffuse struct {
	Texture texture.Texture
}

func (this Material_Diffuse) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Zero
	// return this.Color;
}

func (this Material_Diffuse) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	dir := vec3.Normalize(vec3.Add(h.Normal, RandomDirection()))
	return &ray.Ray{h.Position, dir, 0.00001, math.Inf(1)}
	// return &ray.Ray{ h.Position, h.Normal, 0.00001, math.Inf(1)}
}

func (this Material_Diffuse) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	return this.Texture.SamplePoint(h.SurfaceCoordinates.X, h.SurfaceCoordinates.Y, h.SurfaceCoordinates.Z)
}

func (this Material_Diffuse) String() string { return "Material_Diffuse" }

/*  ======  DIFFUSE_CHECKERBORD  ====== */

type Material_Diffuse_Checkerboard struct {
	Size   float64
	ColorA vec3.Vec3
	ColorB vec3.Vec3
}

func (this Material_Diffuse_Checkerboard) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Zero
}

func (this Material_Diffuse_Checkerboard) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	dir := vec3.Normalize(vec3.Add(h.Normal, RandomDirection()))
	return &ray.Ray{h.Position, dir, 0.000001, math.Inf(1)}
}

func (this Material_Diffuse_Checkerboard) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	x := h.Position.X
	if x < 0 {
		x = math.Abs(x) + 1
	}
	z := h.Position.Z
	if z < 0 {
		z = math.Abs(z) + 1
	}
	if (int(x/this.Size)%2 == 0) != (int(z/this.Size)%2 == 0) {
		return this.ColorA
	} else {
		return this.ColorB
	}
}

func (this Material_Diffuse_Checkerboard) String() string { return "Material_Diffuse_Checkerboard" }

/*  ====== WHITELIGHT ======  */

type Material_WhiteLight struct {
	Emission texture.Texture
	Texture  texture.Texture
}

func (this Material_WhiteLight) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return this.Emission.SamplePoint(h.SurfaceCoordinates.X, h.SurfaceCoordinates.Y, h.SurfaceCoordinates.Z)
}

func (this Material_WhiteLight) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	dir := vec3.Normalize(vec3.Add(h.Normal, RandomDirection()))
	return &ray.Ray{h.Position, dir, 0.000001, math.Inf(1)}
}

func (this Material_WhiteLight) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	return this.Texture.SamplePoint(h.SurfaceCoordinates.X, h.SurfaceCoordinates.Y, h.SurfaceCoordinates.Z)
}

func (this Material_WhiteLight) String() string { return "Material_WhiteLight" }

/*  ====== MIRROR ======  */

type Material_Metal struct {
	Color           vec3.Vec3
	DiffusionFactor float64
}

func (this Material_Metal) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Zero
}

func (this Material_Metal) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	dir := vec3.Subtract(
		r.Direction,
		vec3.Multiply(
			2*vec3.DotProduct(h.Normal, r.Direction),
			h.Normal,
		),
	)

	if this.DiffusionFactor != 0 {
	tryAgain:
		dir = vec3.Normalize(vec3.Add(dir, vec3.Multiply(this.DiffusionFactor, RandomDirection())))
		if vec3.DotProduct(dir, h.Normal) < 0 {
			goto tryAgain
		}
	}

	return &ray.Ray{h.Position, dir, 0.000001, math.Inf(1)}
}

func (this Material_Metal) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	return this.Color
}

func (this Material_Metal) String() string { return "Material_Metal" }

/*  ====== MIRROR_CHECKERBOARD ======  */

type Material_Metal_Checkerboard struct {
	Size            float64
	ColorA          vec3.Vec3
	ColorB          vec3.Vec3
	DiffusionFactor float64
}

func (this Material_Metal_Checkerboard) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Zero
}

func (this Material_Metal_Checkerboard) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	if random.Float64() < 0.8 {
		dir := vec3.Normalize(vec3.Add(h.Normal, RandomDirection()))
		return &ray.Ray{h.Position, dir, 0.000001, math.Inf(1)}
	}
	dir := vec3.Subtract(
		r.Direction,
		vec3.Multiply(
			2*vec3.DotProduct(h.Normal, r.Direction),
			h.Normal,
		),
	)

	if this.DiffusionFactor != 0 {
		dir = vec3.Normalize(vec3.Add(dir, vec3.Multiply(this.DiffusionFactor, RandomDirection())))
	}

	return &ray.Ray{vec3.Add(h.Position, vec3.Zero), dir, 0.000001, math.Inf(1)}
}

func (this Material_Metal_Checkerboard) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	x := h.Position.X
	if x < 0 {
		x = math.Abs(x) + 1
	}
	z := h.Position.Z
	if z < 0 {
		z = math.Abs(z) + 1
	}
	if (int(x/this.Size)%2 == 0) != (int(z/this.Size)%2 == 0) {
		return this.ColorA
	} else {
		return this.ColorB
	}
}

func (this Material_Metal_Checkerboard) String() string { return "Material_Metal_Checkerboard" }

/*  ====== SKY ======  */

type Material_Sky struct {
	Texture texture.Texture
}

func (this Material_Sky) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return this.Texture.SamplePoint(h.SurfaceCoordinates.X, h.SurfaceCoordinates.Y, h.SurfaceCoordinates.Z)
}

func (this Material_Sky) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	return nil
}

func (this Material_Sky) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Black
}

func (this Material_Sky) String() string { return "Material_Sky" }

/*  ====== TRANSPARENT ======  */

type Material_Transparent struct {
	Texture         texture.Texture
	RefractionIndex float64 // air 1.0; water 1.3; glass 1.5;
}

// n1 = outerMaterialRefractionIndex    1.0
// n2 = RefractionIndex                 1.5

func (this Material_Transparent) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Zero
}

func (this Material_Transparent) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	n1 := 1.0
	n2 := this.RefractionIndex
	var rayV ray.Ray
	n := h.Normal

	if vec3.DotProduct(n, r.Direction) > 0 {
		n = vec3.Multiply(-1, n)
		temp := n1
		n1 = n2
		n2 = temp
	}

	c := vec3.DotProduct(vec3.Multiply(-1, n), r.Direction)
	ratio := n1 / n2
	disk := 1 - math.Pow(ratio, 2)*(1-math.Pow(c, 2))
	r0 := math.Pow(((n1 - n2) / (n1 + n2)), 2)
	schlick := r0 + (1-r0)*math.Pow(1+vec3.DotProduct(n, r.Direction), 5)

	if disk >= 0 && random.Float64() > schlick {
		scatDir := vec3.Add(vec3.Multiply(ratio, r.Direction), vec3.Multiply((ratio*c-math.Sqrt(disk)), n))
		rayV = ray.Ray{
			h.Position,
			scatDir,
			0.000001,
			math.Inf(1),
		}
	} else {
		d := vec3.Subtract(r.Direction, vec3.Multiply(2*vec3.DotProduct(n, r.Direction), n))
		rayV = ray.Ray{
			h.Position,
			d,
			0.00001,
			math.Inf(1),
		}
	}
	return &rayV
}

func (this Material_Transparent) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	return this.Texture.SamplePoint(h.SurfaceCoordinates.X, h.SurfaceCoordinates.Y, h.SurfaceCoordinates.Z)
}

func (this Material_Transparent) String() string { return "Material_Transparent" }

/*  ====== NORMAL ======  */
type Material_Normal struct{}

func (this Material_Normal) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Clamp(h.Normal)
	// k := h.T * 20
	// return vec3.Vec3{k,k,k}
}

func (this Material_Normal) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
	return nil
}

func (this Material_Normal) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
	return vec3.Black
}

func (this Material_Normal) String() string { return "Material_Normal" }
