package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func equation() float32 {
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))

	x := int16(rand.Intn(1<<16) - 1 - (1 << 15))
	y := uint8(rand.Intn(1 << 8))
	z := rand.Float32()*(200000.0) - 100000.0

	var s float32
	s = 2*float32(x) + float32(math.Pow(float64(y), 2.0)) - 3/z
	return s
}

func main() {
	fmt.Println("Результат уравнения S = 2 × x + y ^ 2 − 3/z с случайными числами:")
	fmt.Println("S =", equation())
}
