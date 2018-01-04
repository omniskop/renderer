package materials

import (
    "math"
    "customtools/vec3"
    "customtools/space"
    "customtools/ray"
    "cgtools/random"
)

func RandomDirection() vec3.Vec3 {
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

type Material_Diffus2 struct {
    Color       vec3.Vec3
}

func (this Material_Diffus2) EmittedRadiance(r ray.Ray, h space.Hit) vec3.Vec3 {
    return vec3.Zero;
    // return this.Color;
}

func (this Material_Diffus2) ScatteredRay(r ray.Ray, h space.Hit) *ray.Ray {
    dir := vec3.Normalize(vec3.Add(h.Normal, RandomDirection()))
    return &ray.Ray{ h.Position, dir, 0.00001, math.Inf(1)}
    // return &ray.Ray{ h.Position, h.Normal, 0.00001, math.Inf(1)}
}

func (this Material_Diffus2) Albedo(r ray.Ray, h space.Hit) vec3.Vec3 {
    return this.Color
}

func (this Material_Diffus2) String() string {return "Material_Diffus2"}