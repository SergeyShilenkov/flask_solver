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

func (p *Puzzle) commitMove(move *Move) string {
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
		newSnapshot := p.commitMove(move)
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

func (p *Puzzle) snapshot() string {
	snapshot := make([]string, len(p.flasks))
	for idx, flask := range p.flasks {
		snapshot[idx] = flask.String()
	}
	return strings.Join(snapshot, "")
}

func (p *Puzzle) getStageStr() string {
	lines := make([]string, 0, FLASK_SIZE*17)
	line := make([]string, len(p.flasks))

	for i := FLASK_SIZE - 1; i >= 0; i-- {
		for iF, flask := range p.flasks {
			line[iF] = string(flask.balls[i].emoji)
		}
		lines = append(lines, strings.Join(line, ""))
	}

	return strings.Join(lines, "\n")
}

func (p *Puzzle) String() string {
	if !p.Solved {
		return fmt.Sprintf("Просмотрено состояний - %d. Решений нет\n", len(p.states))
	}

	solutionInText := make([]string, 0, 40)

	p.flasks = p.initialFlasks
	solutionInText = append(solutionInText, "Начальная позиция:", p.getStageStr(), "==============================")

	for idx, move := range p.Moves {
		if p.confShowMoves {
			p.makeMove(move)
			solutionInText = append(solutionInText, p.getStageStr())
		}
		solutionInText = append(solutionInText, fmt.Sprintf("%d: %dth %s x%d -> %dth tube", idx, move.from+1, move.verboseRuName, move.ballAmount, move.to+1))
	}

	return fmt.Sprintf("Всего ходов: %d\n%s\n\n", len(p.Moves), strings.Join(solutionInText, "\n"))
}
