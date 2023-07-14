package main

import (
	"fmt"
	// "os"
	// "strconv"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(sumar(3, 5))
	fmt.Println(SumYRest(3, 1))
}

func sumar(n1 int64, n2 int64) int {
	return (int)(n1 + n2)
}

func SumYRest(n1 int64, n2 int64) (int, int) {
	return (int)(n1 + n2), (int)(n1 - n2)
}
