package main

import (
	"fmt"
	"math"
)

//fungsi square akan mengembalikan nilai pangkat 2
//dari dua parameter yang diterima
//contoh
//parameter yang diterima (4,3)
//maka akan mengembalikan 16 dan 9
func main() {

	result1, result2 := square(4, 3)
	fmt.Printf("%d dan %d\n", result1, result2)
	fmt.Println(square(9, 8))
}

//gunakan * untuk melakukan perkalian
// TODO: answer here
func square(number1, number2 int) (int, int) {
	// result1 := number1 * number1
	// result2 := number2 * number2

	// return result1, result2

	result1 := math.Pow(float64(number1), 2) 
	result2 := math.Pow(float64(number2), 2) 

	return int(result1), int(result2)
}