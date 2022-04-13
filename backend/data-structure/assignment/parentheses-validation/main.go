package main

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/parentheses-validation/stack"
)

// Salah satu problem populer yang dapat diselesaikan dengan menggunakan Stack adalah mengecek validitas tanda kurung.
// Diberikan sebuah string yang hanya terdapat karakter '(', ')', '{', '}', '[', dan ']'.
// Tentukan apakah sebuah string merupakan sekumpulan tanda kurung yang valid.
// String tanda kurung yang valid adalah
// 1. Tanda kurung buka harus ditutup oleh pasangannya.
// 2. Tanda kurung buka harus ditutup di urutan yang benar.

// Contoh 1
// Input: s = "()"
// Output: true
// Penjelasan: tanda kurung buka '(' ditutup dengan pasangannya yaitu ')'.

// Contoh 2
// Input: s = "{)"
// Output: false
// Penjelasan: tanda kurung buka '{' tidak ditutup oleh pasangannya.

// Contoh 3
// Input: s = "([])"
// Output: true
// Penjelasan: tanda kurung buka ditutup dengan pasangannya dan sesuai dengan urutan.

func IsValidParentheses(s string) bool {
	// TODO: answer here

	data := []rune{'(', ')', '{', '}', '[', ']'}


	stack := stack.Stack{
		Top:   -1,
		Data: data,
	}

	for _, v := range s {
		if stack.Top == -1 {
			stack.Push(v)
		} else {
			if stack.Top > -1 {
				if stack.Data[stack.Top] == '(' && v == ')' {
					stack.Pop()
					return true
				} else if stack.Data[stack.Top] == '{' && v == '}' {
					stack.Pop()
					return true
				} else if stack.Data[stack.Top] == '[' && v == ']' {
					stack.Pop()
					return true
				} else {
					stack.Push(v)
					return false
				}
			}
		}
	}
	

	return false
}

func main() {
	fmt.Println(IsValidParentheses("(3wew3)"))
	fmt.Println(IsValidParentheses("[{{})]"))
	fmt.Println(IsValidParentheses("([])"))
	fmt.Println(IsValidParentheses("([)]"))
	fmt.Println(IsValidParentheses("{[()]}"))
	fmt.Println(IsValidParentheses("{[(])}"))
}


// parentheseMap := map[rune]rune{
// 	'(': '}',
// 	'{': '}',
// 	'[': ']',
// }

// stack := stack.Stack{
// 	Top:  -1,
// 	Data: []rune{},
// }

// for _, v := range s {
// 	if _, ok := parentheseMap[v]; ok {
// 		stack.Push(v)
// 		return true
// 	} else {
// 		if stack.Top == -1 {
// 			return false
// 		}

// 		if stack.Data[stack.Top] != parentheseMap[v] {
// 			return false
// 		}

// 		stack.Pop()
// 	}
// }