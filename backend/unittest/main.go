package main

import "fmt"

func Sum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Printf("1 + 3 = %d\n", Sum(1, 3))
}
