package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

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

	if s.Top == s.Size - 1 {
		return ErrStackOverflow
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}

	return nil
}
