package main

import (
    "fmt"
    "math"
)

const delta = 1e-4

func Sqrt(x float64, trueValue float64) float64 {
    z := x * 2;
    for math.Abs(trueValue - z) > delta {
        z = z - (z * z - x) / (2 * z)
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2, math.Sqrt(2)))
    fmt.Println(math.Sqrt(2))    
}
