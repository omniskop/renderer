package shapes

import (
	"customtools/space"
	"customtools/vec3"
)

type Torus struct {
	Position vec3.Vec3
	Radius   float64
	Material space.Material
}

// func (s Torus) Intersect(r ray.Ray) *space.Hit {

//     normal := vec3.Divide(vec3.Subtract(point, s.Position), s.Radius);

//     inclination := math.Acos(normal.Y);
//     azimuth := math.Pi + math.Atan2(normal.X, normal.Z);
//     u := azimuth / (2 * math.Pi);
//     v := inclination / math.Pi;

//     return &space.Hit{
//         T: offset,
//         Position: point,
//         Normal: normal,
//         SurfaceCoordinates: vec3.Vec3{u,v,0},
//         Material: s.Material,
//     }
// }
