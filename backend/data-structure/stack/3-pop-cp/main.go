package main

import (
	"errors"
	"fmt"
)

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Size int
	Data []int
}

func NewStack(size int) Stack {
	// TODO: answer here

	stack := Stack{
		Top:  -1,
		Size: size,
		Data: []int{},
	}

	return stack
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if len(s.Data) == s.Size {
		return ErrStackUnderflow
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}

	return nil
}

func (s *Stack) Pop() (int, error) {
	// TODO: answer here

	if s.Top == -1 {
		return 0, ErrStackUnderflow
	}

	poppedValue := s.Data[s.Top]
	s.Top -= 1
	s.Data = s.Data[:len(s.Data)-1]
	return poppedValue, nil
}

func main() {
	stack := NewStack(10)

	for i := 0; i < stack.Size; i++ {
		stack.Top += 1
		stack.Data = append(stack.Data, i)
	}
	fmt.Println(stack.Data)


	pop, _ := stack.Pop()
	fmt.Println(pop)

}
