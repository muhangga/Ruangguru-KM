package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

type Stack struct {
	// TODO: answer here

	Top int
	Size int
	Data []int

}

func NewStack(size int) Stack {
	// TODO: answer here

	stack := Stack{
		Top: -1,
		Size: size,
		Data: []int{},
	}

	return stack
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here

	if s.Top == s.Size -1 {
		return ErrEmptyStack
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}

	return nil
}

func (s *Stack) Peek() (int, error) {
	// TODO: answer here

	if s.Top == -1 {
		return 0, ErrEmptyStack
	}

	return s.Data[s.Top], nil

}
