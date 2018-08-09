package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func main() {
	fmt.Printf("1 + 2 = %d\n", add(1, 2))
	fmt.Printf("1 - 2 = %d\n", minus(1, 2))
}
