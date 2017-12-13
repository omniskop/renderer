package shapes

import (
    "customtools/ray"
    "customtools/space"
)

type Shape interface {
    Intersect(r ray.Ray) *space.Hit
}