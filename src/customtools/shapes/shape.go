package shapes

import (
    "customtools/hit"
    "customtools/ray"
)

type shape interface {
    Intersect(r ray.Ray) *hit.Hit
}