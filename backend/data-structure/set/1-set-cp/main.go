// Tulis program sebagai fungsi yang memeriksa apakah dua set dikatakan intersection atau tidak.
// Jika kedua set intersection, fungsi tersebut harus mengembalikan nilai intersection nya.
//
// Contoh 1
// Input: a = {"Java", "Python", "Javascript", "C ++", "C#"}, b = {"Swift", "Java", "Kotlin", "Python"}
// Output: {"Java", "Python"}
// Explanation: intersection dari a dan b adalah "Java" dan "Python"
//
// Contoh 2
// Input: a = {"Java", "Python"}, b = {"Swift"}
// Output: {}
// Explanation: tidak ada intersection dari a dan b

package main

import "fmt"

func main() {
	var str1 = []string{"Java", "Python", "Javascript", "C ++", "C#"}
	var str2 = []string{"Swift", "Java", "Kotlin", "Python"}
	fmt.Println(Intersection(str1, str2))
}

func Intersection(str1, str2 []string) (inter []string) {

	hash := make(map[string]bool)

	for _, item := range str1 {
		hash[item] = true
	}

	for _, item := range str2 {
		if _, exist := hash[item]; exist {
			inter = append(inter, item)
		}
	}

	inter = RemoveDuplicates(inter)

	return inter
}

func RemoveDuplicates(elements []string) (nodups []string) {

	encountered := make(map[string]bool)

	for _, v := range elements {
		if encountered[v] == true {
			continue
		} else {
			// Add element to new slice.
			nodups = append(nodups, v)
			// Remember that we have encountered this element.
			encountered[v] = true
		}
	}

	return nodups
}
