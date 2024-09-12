package validator

import "fmt"

type AmountEmptySlots struct {
	Amount int
}

func (e *AmountEmptySlots) Error() string {
	return fmt.Sprintf("There are %d empty slots, but 8 is required", e.Amount)
}

type AmountFlasks struct {
	Amount int
}

func (e *AmountFlasks) Error() string {
	return fmt.Sprintf("There are %d flasks, but 6 is min", e.Amount)
}

type AmountSpecificColorError struct {
	Color     string
	Amount    int
	maxAmount int
}

func (e *AmountSpecificColorError) Error() string {
	return fmt.Sprintf("There are %d %s, but %d is max", e.Amount, e.Color, e.maxAmount)
}

type UnknownColor struct {
	Color string
}

func (e *UnknownColor) Error() string {
	return fmt.Sprintf("%s is unknown", e.Color)
}

type AmountLinesError struct {
	CurrentAmount  int
	ExpectedAmount int
}

func (e *AmountLinesError) Error() string {
	return fmt.Sprintf("There are %d lines, but %d expected", e.CurrentAmount, e.ExpectedAmount)
}
