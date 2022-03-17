package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
	var jari, tinggi, volume float64
	pi := 3.14

	fmt.Printf("Masukan jari-jari :")
	fmt.Scan(&jari)

	fmt.Printf("Masukan tinggi :")
	fmt.Scan(&tinggi)

	volume = pi * jari * jari *tinggi
	fmt.Printf("Jadi volumenya adalah: %.4f", volume)
}
