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

type Material_Diffuse struct {
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

type Material_WhiteLight struct {
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

/*  ====== MIRROR ======  */

type Material_Mirror struct {
    Color   vec3.Vec3
}

func (this Material_Mirror) EmittedRadiance(r ray.Ray, h Hit) vec3.Vec3 {
    return vec3.Zero
}

func (this Material_Mirror) ScatteredRay(r ray.Ray, h Hit) *ray.Ray {
    n := h.Normal
    p := h.Position
    P := vec3.Subtract(p, r.Direction)
    
    t := (vec3.DotProduct(n, P) - n.X*p.X - n.Y*p.Y - n.Z*p.Z) / vec3.DotProduct(n,n)
    
    
    M := vec3.Add(p, vec3.Multiply(t, n))
    
    
    Q := vec3.Add(P, vec3.Multiply(2,  vec3.Subtract(M, P)   ))
    
    N := vec3.Normalize(vec3.Subtract(Q, p))
    
    return &ray.Ray{ vec3.Add(h.Position, vec3.Divide(N, 1000)), N, 0.000001, math.Inf(1)}
}

func (this Material_Mirror) Albedo(r ray.Ray, h Hit) vec3.Vec3 {
    return this.Color
}

/*  ====== SKY ======  */

type Material_Sky struct {
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