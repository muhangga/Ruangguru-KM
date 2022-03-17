package main

import (
	"fmt"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here
	var choice int
	var result float32

	fmt.Println("1 : Rectange area")
	fmt.Println("2 : Rectangular area")
	fmt.Println("3 : Triangle area")
	fmt.Println("4 : Circle area")

	fmt.Printf("Masukan pilihan :")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		var sisi float32

		fmt.Printf("Masukan sisi :")
		fmt.Scan(&sisi)

		result = sisi * sisi
		fmt.Printf("Luas Persegi adalah: %f", result)
		break
	case 2:
		var p, l float32

		fmt.Printf("Masukan panjang :")
		fmt.Scan(&p)

		fmt.Printf("Masukan lebar :")
		fmt.Scan(&l)

		result = p * l
		fmt.Printf("Luas Persegi Panjang adalah: %f", result)
		break

	case 3:
		var a, t float32

		fmt.Printf("Masukan alas :")
		fmt.Scan(&a)

		fmt.Printf("Masukan tinggi :")
		fmt.Scan(&t)

		result = 0.5 * a * t
		fmt.Printf("Luas Persegi Panjang adalah: %f", result)
		break
	case 4:
		var r float32
		const pi = 3.14

		fmt.Printf("Masukan jari-jari :")
		fmt.Scan(&r)

		result = pi * r * r
		fmt.Printf("Luas lingkaran adalah: %f", result)
		break
	default:
		fmt.Println("Wrong choice")

	}
}
