package main

import (
	"flag"
	"flask_solver/parsers"
	"flask_solver/solver"
	"flask_solver/validator"
	"fmt"
	"os"
	"time"
)

const (
	FLASK_SIZE = 4
)

var maxGoroutines = flag.Int("g", 50, "maximum amount of goroutines")
var pathDataFlasks = flag.String("p", "./PrideNN/lvl1.txt", "path to file with flasks")
var numbUnknownColors = flag.Int("NUC", 0, "number of unknown colors to discover")
var unknownColorIsLast = flag.Bool("UCL", false, "atleast one unknown color is last")
var usePermutations = flag.Bool("pr", false, "use permutations to solve a puzzle")
var showMoves = flag.Bool("m", false, "show moves in text (false) or stages (true)")

func main() {
	flag.Parse()

	// Парсинг данных
	p, err := parsers.NewParser(pathDataFlasks)
	if err != nil {
		fmt.Printf("ОШИБКА в parsers.NewParser(pathDataFlasks)!\n%s", err)
		os.Exit(1)
	}
	parsedData, err := (*p).Parse()
	if err != nil {
		fmt.Printf("ОШИБКА в (*p).Parse()!\n%s", err)
		os.Exit(1)
	}

	// Валидация полученных данных
	err = validator.ValidateData(&parsedData)
	if err != nil {
		fmt.Printf("ОШИБКА в validator.ValidateData(parsedData)!\n%s", err)
		os.Exit(1)
	}

	// Решение задачи
	start := time.Now()

	solution, err := solver.SolvePuzzle(&solver.ConfigData{
		MaxGoroutines:      *maxGoroutines,
		NumbUnknownColors:  *numbUnknownColors,
		UnknownColorIsLast: *unknownColorIsLast,
		UsePermutations:    *usePermutations,
		ShowMoves:          *showMoves}, &parsedData)

	elapsedTime := time.Since(start)
	if err != nil {
		fmt.Printf("ОШИБКА в solver.SolvePuzzle(&solver.ConfigData{})!\n%s", err)
		os.Exit(1)
	}

	// Вывод результатов
	fmt.Printf("Выполнено за %s\n%s", elapsedTime, solution)
}
