package main

import (
	"fmt"
	"math"
)

func main() {
	f := func(x, y, z int) float32 { return 2*float32(x) + float32(math.Pow(float64(y), 2.0)) - 3/float32(z) }
	fmt.Println("Результат уравнения S = 2 × x + y ^ 2 − 3/z:")
	fmt.Println("S =", f(10, 20, 30))
}
