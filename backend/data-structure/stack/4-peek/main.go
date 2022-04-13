package main

import (
	"errors"
	"fmt"
)

// Melihat elemen teratas pada stack disebut dengan peek.

type Stack struct {
	Top  int
	Data []int
}

func (s *Stack) Peek() (int, error) {
	if s.Top == -1 {
		return 0, errors.New("stack is empty")
	} else {
		return s.Data[s.Top], nil
	}
}


func main() {
	stack := Stack{
		Top:  -1,
		Data: []int{},
	}

	stack.Top += 1
	stack.Data = append(stack.Data, 1)
	stack.Top += 1
	stack.Data = append(stack.Data, 2)
	stack.Top += 1
	stack.Data = append(stack.Data, 3)

	fmt.Println(stack.Peek())
}