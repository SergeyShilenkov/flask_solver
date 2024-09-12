package tparser

import "fmt"

type LengthFileError struct {
	Length int
}

func (e *LengthFileError) Error() string {
	return fmt.Sprintf("There are more than 4 lines in file (%d)", e.Length)
}

type LengthLineError struct {
	Line           int
	Amount         int
	ExpectedAmount int
}

func (e *LengthLineError) Error() string {
	return fmt.Sprintf("There are %d elements in %d, but %d expected", e.Amount, e.Line, e.ExpectedAmount)
}

type UnknownColor struct {
	Color string
}

func (e *UnknownColor) Error() string {
	return fmt.Sprintf("%s is unknown", e.Color)
}
