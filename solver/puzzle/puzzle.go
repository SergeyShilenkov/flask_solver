package puzzle

import (
	"fmt"
	"strings"
)

const (
	FLASK_SIZE = 4
)

type Puzzle struct {
	confNumbUnknownColors  int
	confUnknownColorIsLast bool
	confShowMoves          bool
	initialFlasks          []*Flask
	flasks                 []*Flask
	Moves                  []*Move
	states                 map[string]struct{}
	Solved                 bool
}

func NewPuzzle(initialColors [][FLASK_SIZE]int, uc []int, numbUnknownColors int, unknownColorIsLast bool, showMoves bool) *Puzzle {
	flasks := make([]*Flask, len(initialColors))
	initialStage := make([]*Flask, len(initialColors))
	idxUc := 0

	for idx, flask := range initialColors {
		for idxC, c := range flask {
			if c == 1 && uc != nil {
				flask[idxC] = uc[idxUc]
				idxUc++
			} else {
				flask[idxC] = c
			}
		}
		flasks[idx] = NewFlask(idx, flask)
		initialStage[idx] = NewFlask(idx, flask)
	}

	return &Puzzle{
		confNumbUnknownColors:  numbUnknownColors,
		confUnknownColorIsLast: unknownColorIsLast,
		confShowMoves:          showMoves,
		initialFlasks:          initialStage,
		flasks:                 flasks,
		Moves:                  []*Move{},
		states:                 make(map[string]struct{}),
		Solved:                 false,
	}
}

func (p *Puzzle) isMaxUnknownColors() bool {
	if p.confNumbUnknownColors <= 0 {
		return false
	}

	numb := 0
	unknownBallIsLast := false

	for _, flask := range p.flasks {
		if flask.isEmpty() {
			continue
		}

		upperBalls := flask.upperBalls()
		if upperBalls.color == UNKNOWN {
			if p.confUnknownColorIsLast && (upperBalls.amount == 1) {
				unknownBallIsLast = true
			}
			numb++
		}
	}

	return (numb >= p.confNumbUnknownColors) && (p.confUnknownColorIsLast && unknownBallIsLast || !p.confUnknownColorIsLast)
}

func (p *Puzzle) isSolved() bool {
	result := true
	for _, flask := range p.flasks {
		if !flask.isSolved() {
			result = false
			break
		}
	}

	return result || p.isMaxUnknownColors()
}

func (p *Puzzle) getPossibleMoves() []*Move {
	moves := make([]*Move, 0, 10)

	for idx_f, flask := range p.flasks {
		if flask.isSolved() || flask.isAlmostSolved() {
			continue
		}

		upperballs := flask.upperBalls()
		if upperballs.color == UNKNOWN {
			continue
		}

		for idx_pf, potentional_flask := range p.flasks {
			if idx_pf == idx_f || (flask.hasOneColor() && potentional_flask.isEmpty()) {
				continue
			}
			if potentional_flask.canReceive(upperballs.color) {
				upperballs.amount = min(upperballs.amount, potentional_flask.freeSpace())

				moves = append(moves, NewMove(flask.num, potentional_flask.num, upperballs))
			}
		}
	}

	return moves
}

func (p *Puzzle) makeMove(move *Move) {
	ball := p.flasks[move.from].pop(move.ballAmount)
	p.flasks[move.to].push(ball, move.ballAmount)
}

func (p *Puzzle) commitMove(move *Move) []byte {
	p.makeMove(move)
	p.Moves = append(p.Moves, move)
	return p.snapshot()
}

func (p *Puzzle) rollBackMove(move *Move) {
	ball := p.flasks[move.to].pop(move.ballAmount)
	p.flasks[move.from].push(ball, move.ballAmount)
	p.Moves = p.Moves[:len(p.Moves)-1]
}

func (p *Puzzle) Solve() bool {
	if p.isSolved() {
		p.Solved = true
		return true
	}

	for _, move := range p.getPossibleMoves() {
		newSnapshot := string(p.commitMove(move))
		_, exist := p.states[newSnapshot]
		if exist {
			p.rollBackMove(move)
			continue
		}

		p.states[newSnapshot] = struct{}{}
		if p.Solve() {
			return true
		}
		p.rollBackMove(move)
	}

	return false
}

func (p *Puzzle) snapshot() []byte {
	snapshot := make([]byte, 0, len(p.flasks)*FLASK_SIZE*2)
	for _, flask := range p.flasks {
		snapshot = append(snapshot, flask.String()...)
	}

	return snapshot
}

func (p *Puzzle) getStageStr() []byte {
	stage := make([]byte, 0, len(p.flasks)*FLASK_SIZE*8)

	for i := FLASK_SIZE - 1; i >= 0; i-- {
		for _, flask := range p.flasks {
			stage = append(stage, []byte(flask.balls[i].emoji)...)
		}
		if i != 0 {
			stage = append(stage, '\n')
		}
	}

	return stage
}

func (p *Puzzle) getColorMaxLength() int {
	length := -1 << 32

	for _, f := range p.flasks {
		for _, c := range f.balls {
			if c != UNKNOWN && c != EMPTY && length < len([]rune(c.verbose)) {
				length = len([]rune(c.verbose))
			}
		}
	}

	return length
}

func (p *Puzzle) String() string {
	if !p.Solved {
		return fmt.Sprintf("Просмотрено состояний - %d. Решений нет\n", len(p.states))
	}

	solutionInText := make([]string, 0, 40)
	lengthColor := p.getColorMaxLength()

	p.flasks = p.initialFlasks
	solutionInText = append(solutionInText, "Начальная позиция:", string(p.getStageStr()), "==============================")

	for idx, move := range p.Moves {
		if p.confShowMoves {
			p.makeMove(move)
			solutionInText = append(solutionInText, string(p.getStageStr()), fmt.Sprintf("%2d: %2dth %s x%d -> %dth tube", idx+1, move.from+1, move.verbose, move.ballAmount, move.to+1))
			continue
		}
		solutionInText = append(solutionInText, fmt.Sprintf("%2[1]d: %[2]*[3]s x%[4]d %2[5]dth -> %[6]dth tube", idx+1, lengthColor, move.verbose, move.ballAmount, move.from+1, move.to+1))
	}

	return fmt.Sprintf("Всего ходов: %d\n%s\n\n", len(p.Moves), strings.Join(solutionInText, "\n"))
}
