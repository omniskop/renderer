package shapes

import(
    "cgtools/random"
    "customtools/vec3"
    "customtools/ray"
    "math"
)

func RandomCosDirection() vec3.Vec3 {
    direction := vec3.One
    for direction.Length() > 1 {
        direction = vec3.Vec3{
            random.Float64() * 2 - 1,
            random.Float64() * 2 - 1,
            random.Float64() * 2 - 1,
        }
    }
    return direction
}

type Material interface {
    EmittedRadiance(r ray.Ray, h Hit)       vec3.Vec3
    ScatteredRay(r ray.Ray, h Hit)          *ray.Ray
    Albedo(r ray.Ray, h Hit)                vec3.Vec3
}

/*  ======  DIFFUSE  ====== */

type Material_Diffuse struct { // Lambertsches Material
    Color       vec3.Vec3
}

func (this Material_Diffuse) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
    return vec3.Zero;
}

func (this Material_Diffuse) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
    dir := vec3.Normalize(vec3.Add(h.Normal, RandomCosDirection()))
    return &ray.Ray{ h.Position, dir, 0.000001, math.Inf(1)}
}

func (this Material_Diffuse) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
    return this.Color
}

/*  ====== WHITELIGHT ======  */

type Material_WhiteLight struct { // Lambertsches Material
    Color       vec3.Vec3
}

func (this Material_WhiteLight) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
    return vec3.White
}

func (this Material_WhiteLight) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
    dir := vec3.Normalize(vec3.Add(h.Normal, RandomCosDirection()))
    return &ray.Ray{h.Position, dir, 0.000001, math.Inf(1)}
}

func (this Material_WhiteLight) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
    return this.Color
}

/*  ====== SKY ======  */

type Material_Sky struct { // Lambertsches Material
    Color       vec3.Vec3
}

func (this Material_Sky) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
    return this.Color
}

func (this Material_Sky) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
    return nil
}

func (this Material_Sky) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
    return vec3.Black
}