package solver

import "fmt"

type IncorrectAmountGoroutines struct {
	Amount int
	Min    int
	Max    int
}

func (e *IncorrectAmountGoroutines) Error() string {
	return fmt.Sprintf("There are %d goroutines, but amount should be in %d-%d", e.Amount, e.Min, e.Max)
}

type UndefinedColors struct {
	Amount    int
	FlaskSize int
}

func (e *UndefinedColors) Error() string {
	return fmt.Sprintf("There are undefined colors in data. %d colors", e.Amount/e.FlaskSize)
}

type NoUnknownColors struct{}

func (e *NoUnknownColors) Error() string {
	return "There are no unknown colors"
}

type TooManyUnknownColors struct {
	Amount int
}

func (e *TooManyUnknownColors) Error() string {
	return fmt.Sprintf("There are too many unknown colors - %d", e.Amount)
}
