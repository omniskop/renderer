package random

import (
    "math/rand"
    "time"
)

func InitSeed() { 
    rand.Seed(time.Now().UnixNano())
}

func IntRange(min, max int) int {
    return rand.Intn(max - min) + min
}

func Float64() float64 {
    return rand.Float64()
}