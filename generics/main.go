package main

import "fmt"

func add[T int | float64](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add(1.3, 2))
}
