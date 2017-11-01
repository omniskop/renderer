package hit

import (
    "customtools/vec3"
    "customtools/ray"
)

type Hit struct {
    T           ray.Ray
    Position    vec3.Vec3
    Normal      vec3.Vec3
}