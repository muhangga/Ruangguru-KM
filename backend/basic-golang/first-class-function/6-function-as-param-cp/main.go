package main

import "fmt"

func main() {
	//fungsi untuk menampilkan string dan memiliki parameter fungsi
	//fungsi yang akan dimasukkan adalah fungsi yang memberi selamat malam
	// TODO: answer here
	printString := func(f func () string) {
		fmt.Println(f())	
	}

	goodNight := func () string  {
		return "Selamat Malam"
	}

	printString(goodNight)

}
