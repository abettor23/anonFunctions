package main

import (
	"fmt"
)

var count int = 0

func wrapping(A func(int, int) int) {
	defer fmt.Println(A(10, 5))
	fmt.Printf("Расчет %d \n", count+1)
}

func main() {

	finalScore := 3
	superFunc := func(x, y int) int {
		switch count {
		case 1:
			return x - y
		case 2:
			return x * y
		default:
			return x + y
		}
	}

	for count < finalScore {
		wrapping(superFunc)
		count++
	}
}
