package hit

import (
    "cgtools/vec3"
    "cgtools/ray"
)

type Hit struct {
    T           ray.Ray
    Position    vec3.Vec3
    Normal      vec3.Vec3
}