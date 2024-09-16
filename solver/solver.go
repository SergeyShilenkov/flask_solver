package solver

import (
	"flask_solver/solver/puzzle"
	"fmt"
	"math"

	"gonum.org/v1/gonum/stat/combin"
)

type ConfigData struct {
	MaxGoroutines      int
	NumbUnknownColors  int
	UnknownColorIsLast bool
	UsePermutations    bool
	ShowMoves          bool
}

func colorsInHash(colors []int, permutation []int) int64 {
	result := int64(0)

	for i := 0; i < len(colors); i++ {
		result += int64(math.Pow(float64(100), float64(len(permutation)-1-i))) * int64(colors[permutation[i]])
	}

	return result
}

func getUnknownColors(data *[][puzzle.FLASK_SIZE]int) (int, []int, error) {
	numbUnknownColors := make(map[int]int)
	numbUnknownColors[1] = (len(*data) - 2) * puzzle.FLASK_SIZE
	totalNumbUnknownColors := numbUnknownColors[1]

	for _, f := range *data {
		for _, c := range f {
			if c != 0 && c != 1 {
				if _, ok := numbUnknownColors[c]; ok {
					if numbUnknownColors[c] > 1 {
						numbUnknownColors[c]--
					} else {
						delete(numbUnknownColors, c)
					}
				} else {
					numbUnknownColors[c] = puzzle.FLASK_SIZE - 1
					numbUnknownColors[1] -= puzzle.FLASK_SIZE
				}
				totalNumbUnknownColors--
			}
		}
	}
	if numbUnknownColors[1] >= puzzle.FLASK_SIZE {
		return 0, nil, &UndefinedColors{Amount: numbUnknownColors[1], FlaskSize: puzzle.FLASK_SIZE}
	}
	if len(numbUnknownColors) == 1 {
		return 0, nil, &NoUnknownColors{}
	}
	if len(numbUnknownColors)-1 > 8 {
		return 0, nil, &TooManyUnknownColors{Amount: len(numbUnknownColors) - 1}
	}

	unknownColors := make([]int, totalNumbUnknownColors)
	idx := 0
	for k, v := range numbUnknownColors {
		if k != 1 {
			for i := 0; i < v; i++ {
				unknownColors[idx] = k
				idx++
			}
		}
	}

	return totalNumbUnknownColors, unknownColors, nil
}

func parallelSolvePuzzles(data [][puzzle.FLASK_SIZE]int, maxGoroutines int, showMoves bool) (string, error) {
	totalNumbUnknownColors, unknownColors, err := getUnknownColors(&data)
	if err != nil {
		return "", err
	}

	permutations := combin.Permutations(totalNumbUnknownColors, totalNumbUnknownColors)

	ch := make(chan *puzzle.Puzzle, maxGoroutines)
	hashPuzzles := make(map[int64]struct{})

	var intPuzzle *puzzle.Puzzle
	amountMoves := math.MaxInt32
	amountUnsolvedPermutations := 0
	amountRepeatPermutations := 0

	sp := 0

	for sp < len(permutations) {
		ip := 0
		for ip < maxGoroutines && (sp+ip < len(permutations)) {
			cis := colorsInHash(unknownColors, permutations[sp+ip])
			if _, exist := hashPuzzles[cis]; exist {
				amountRepeatPermutations++
				sp++
				continue
			}

			hashPuzzles[cis] = struct{}{}

			go func(n int) {
				for i := 0; i < len(unknownColors); i++ {
					permutations[n][i] = unknownColors[permutations[n][i]]
				}

				p := puzzle.NewPuzzle(data, permutations[n], 0, false, showMoves)
				p.Solve()

				ch <- p
			}(sp + ip)
			ip++
		}

		for ipd := 0; ipd < ip; ipd++ {
			result := <-ch
			if result.Solved {
				if amountMoves > len(result.Moves) {
					amountMoves = len(result.Moves)
					intPuzzle = result
				}
			} else {
				amountUnsolvedPermutations++
			}
		}

		sp += ip
	}
	close(ch)

	uColorsEmoji := make([]*puzzle.Color, len(unknownColors))
	for idx, c := range unknownColors {
		uColorsEmoji[idx] = puzzle.COLORCONVERT[c]
	}

	return fmt.Sprintf("ПРОВЕРЕНО: %d\nУНИКАЛЬНЫХ ВАРИАНТОВ: %d\nНевозможных комбинаций: %d\nПовторных комбинаций: %d\nНеизвестные цвета: %v\n%s", sp, len(hashPuzzles), amountUnsolvedPermutations, amountRepeatPermutations, uColorsEmoji, intPuzzle.String()), nil
}

func SolvePuzzle(config *ConfigData, data *[][puzzle.FLASK_SIZE]int) (string, error) {
	if config.UsePermutations {
		if config.MaxGoroutines < 2 || config.MaxGoroutines > 500 {
			return "", &IncorrectAmountGoroutines{Amount: config.MaxGoroutines, Min: 2, Max: 500}
		}
		solution, err := parallelSolvePuzzles(*data, config.MaxGoroutines, config.ShowMoves)
		if err != nil {
			return "", err
		}
		return solution, nil
	} else {
		task := puzzle.NewPuzzle(*data, nil, config.NumbUnknownColors, config.UnknownColorIsLast, config.ShowMoves)
		task.Solve()
		return task.String(), nil

	}
}
