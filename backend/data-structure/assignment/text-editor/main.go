package main

import (
	// "fmt"

	"github.com/ruang-guru/playground/backend/data-structure/assignment/text-editor/stack"
)

// Implementasikan teks editor sederhana
// Teks editor ini memiliki 4 method
// Write: digunakan untuk menulis per karakter
// Read: digunakan untuk mencetak karakter yang telah ditulis
// Undo: digunakan untuk melakukan operasi undo
// Redo: digunakan untuk melakukan operasi redo

type TextEditor struct {
	UndoStack *stack.Stack
	RedoStack *stack.Stack
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		UndoStack: stack.NewStack(),
		RedoStack: stack.NewStack(),
	}
}

func (te *TextEditor) Write(ch rune) {
	// TODO: answer here
	te.UndoStack.Push(ch)
	te.RedoStack.Pop()
}

func (te *TextEditor) Read() []rune {
	// TODO: answer here

	var result []rune
	data := te.UndoStack.Data

	for _, v := range data {
		result = append(result, v)
	}
	return result
}

func (te *TextEditor) Undo() {
	checkUndo, err := te.UndoStack.Peek()
	if err != nil {
		return
	}

	te.UndoStack.Pop()
	te.RedoStack.Push(checkUndo)
}

func (te *TextEditor) Redo() {
	// TODO: answer here

	checkRedo, err := te.RedoStack.Peek()
	if err != nil {
		return
	}

	te.RedoStack.Pop()
	te.UndoStack.Push(checkRedo)
}
