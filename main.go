package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("1 + 2 = %d\n", add(1, 2))
}
