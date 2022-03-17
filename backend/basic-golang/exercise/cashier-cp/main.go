package main

import (
	"fmt"
)

// Check Point:
// Cashier
// - Input: Nama Menu, Add Again
// - Output: Total Harga yang harus dibayar

// Contoh:
// Menu makanan:
// 1 . Menu: bakso ,  Price : 20000
// 2 . Menu: burger ,  Price : 15000
// 3 . Menu: sate ,  Price : 25000
// 4 . Menu: sosis ,  Price : 20000
// 5 . Menu: soto ,  Price : 25000

// Input:
// - Form:
//   - Masukan nama menu yang mau dipesan: bakso

// Output:
// Menu telah ditambahkan ke Ordered Menu:
// Menu: bakso ,  Price : 20000

// Input:
// - Ingin memesan menu lain?(yes/no): if no (break) --> if yes (add again)

// Output:
// Data Order Updated:
//  - Menu: sate ,  Price : 25000
//  - Menu: burger ,  Price : 15000
//  - Menu: sosis ,  Price : 20000
//  - Total harga makanan yang harus anda bayar:  60000

func main() {
	foodMenu := map[string]int64{
		"bakso":  20000,
		"burger": 15000,
		"sate":   25000,
		"sosis":  20000,
		"soto":   25000,
	}

	orderMenu := make(map[string]int64)

	// TODO: answer here
	// var menu, addAgain string
	for {
		fmt.Println("Menu Makanan")
		for key, val := range foodMenu {
			fmt.Println("- Menu :", key, ", Price : ", val)
		}

		var menu, addAgain string
		fmt.Println()

		for {
			fmt.Print("Masukan nama menu : ")
			fmt.Scan(&menu)

			if price, available := foodMenu[menu]; available {
				orderMenu[menu] = price
				break
			} else {
				fmt.Printf("%v menu tidak di temukan, silahkan coba lagi\n", menu)
			}
		}

		fmt.Println()
		fmt.Printf("Menu telah ditambahkan ke Order Menu")
		for menu, price := range foodMenu {
			fmt.Println(" - Menu:", menu, ", ", "Price :", price)
		}

		fmt.Printf("Ingin menambah pesanan baru? (y/n) : ")
		fmt.Scan(&addAgain)

		if addAgain == "n" {
			fmt.Printf("Data Order Terupdate")
			for key, val := range foodMenu {
				fmt.Println("- Menu :", key, ", Price : ", val)
			}
			break
		}
	}
}
