package shapes

import (
    "customtools/hit"
    "customtools/ray"
)

type Shape interface {
    Intersect(r ray.Ray) *hit.Hit
}