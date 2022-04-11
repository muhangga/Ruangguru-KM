// Mengecek jika dua string adalah anagram
// Anagram adalah kata yang dibentuk melalui penataan ulang huruf-huruf dari beberapa kata lain.
//
// Contoh 1
// Input: a = "keen", b = "knee"
// Output: "Anagram"
// Explanation: Jika ditata, "knee" dan "keen" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 2
// Input: a = "fried", b = "fired"
// Output: "Anagram"
// Explanation: Jika ditata, "fried" dan "fired" memiliki huruf-huruf yang sama, hanya berbeda urutan
//
// Contoh 3
// Input: a = "apple", b = "paddle"
// Output: "Bukan Anagram"
// Explanation: Jika ditata, "apple" dan "paddle" memiliki huruf-huruf yang berbeda

package main

import "fmt"

func main() {
	var str1 = "fried"
	var str2 = "fired"
	fmt.Println(AnagramsChecker(str1, str2))
}

func AnagramsChecker(str1 string, str2 string) string {

	if len(str1) != len(str2) {
		return "Bukan Anagram"
	}

	hash := make(map[string]int)

	for _, r := range str1 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	for _, r := range str2 {
		k := hash[string(r)]

		if k == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = k - 1
		}
	}

	var isAnagram = true
	for _, value := range hash {
		if value != 0 {
			isAnagram = false
			break
		}
	}

	if isAnagram {
		return "Anagram"
	} else {
		return "Bukan Anagram"
	}

	// for _, val := range str1 {
	// 	str1Map[val]++

	// }

	// for _, val := range str2 {
	// 	str2Map[val]++
	// }

	// for key, val := range str1Map {
	// 	if val != str2Map[key] {
	// 		return "Bukan Anagram"
	// 	}
	// }

	// return "Anagram"
}
