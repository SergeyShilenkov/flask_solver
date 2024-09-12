package tparser

import (
	"bufio"
	"os"
	"strings"
)

var convertMap = map[string]int{
	"-":              0,
	"?":              1,
	"бирюзовый":      2,
	"жёлтый":         3,
	"зелёный":        4,
	"коричневый":     5,
	"красный":        6,
	"малиновый":      7,
	"оранжевый":      8,
	"розовый":        9,
	"светло-зелёный": 10,
	"серый":          11,
	"синий":          12,
	"сиреневый":      13,
	"тёмно-синий":    14,
	"фиолетовый":     15,
}

type TextParser struct {
	path *string
}

func New(path *string) *TextParser {
	return &TextParser{path: path}
}

func (p *TextParser) Parse() ([][4]int, error) {
	file, err := os.Open(*p.path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, strings.TrimSpace(fileScanner.Text()))
	}
	file.Close()

	if len(fileLines) > 4 {
		return nil, &LengthFileError{Length: len(fileLines)}
	}

	lenFirstLine := len(strings.Split(fileLines[0], ","))
	parsedData := make([][4]int, lenFirstLine)

	for idx, line := range fileLines {
		parsedLine := strings.Split(line, ",")
		if len(parsedLine) != lenFirstLine {
			return nil, &LengthLineError{Line: idx, Amount: len(parsedLine), ExpectedAmount: lenFirstLine}
		}

		for idx_c, c := range parsedLine {
			cv, ok := convertMap[strings.TrimSpace(c)]
			if !ok {
				return nil, &UnknownColor{Color: strings.TrimSpace(c)}
			}
			parsedData[idx_c][4-idx-1] = cv
		}
	}

	return parsedData, nil
}
