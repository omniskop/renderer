package circle

import (
    "cgtools/vec3"
    "cgtools/random"
    "sort"
)

type Circle struct {
    Pos     position
    Color   vec3.Vec3
    Radius  int
}

type position struct {
    X   int
    Y   int
}

func NewRandom(width, height int) Circle {
    return Circle{
        position{ random.IntRange(0, width), random.IntRange(0, height) },
        vec3.Vec3{ random.Float64(), random.Float64(), random.Float64() },
        random.IntRange(height / 20, height / 3),
    }
}

func SortByRadius(circles []Circle) {
    sort.Sort(byRadius(circles))
}

type byRadius []Circle

func (s byRadius) Len() int {
    return len(s)
}
func (s byRadius) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s byRadius) Less(i, j int) bool {
    return s[i].Radius < s[j].Radius
}