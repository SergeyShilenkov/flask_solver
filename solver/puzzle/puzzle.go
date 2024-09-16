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
	tmpFlask := make([]int, FLASK_SIZE)
	idxUc := 0

	for idx, flask := range initialColors {
		for idxC, c := range flask {
			if c == 1 && uc != nil {
				tmpFlask[idxC] = uc[idxUc]
				idxUc++
			} else {
				tmpFlask[idxC] = c
			}
		}
		flasks[idx] = NewFlask(idx, tmpFlask)
		initialStage[idx] = NewFlask(idx, tmpFlask)
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
	moves := []*Move{}

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
				if upperballs.amount > potentional_flask.freeSpace() {
					upperballs.amount = potentional_flask.freeSpace()
				}
				moves = append(moves, NewMove(flask.num, potentional_flask.num, upperballs))
			}
		}

	}
	return moves
}

func (p *Puzzle) makeMove(move *Move) {
	for i := 0; i < move.ballAmount; i++ {
		ball := p.flasks[move.from].pop()
		p.flasks[move.to].push(ball)
	}
}

func (p *Puzzle) commitMove(move *Move) string {
	p.makeMove(move)
	p.Moves = append(p.Moves, move)
	return p.snapshot()
}

func (p *Puzzle) rollBackMove(move *Move) {
	for i := 0; i < move.ballAmount; i++ {
		ball := p.flasks[move.to].pop()
		p.flasks[move.from].push(ball)
	}
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
	lines := make([]string, 0, 4*17)

	for line := FLASK_SIZE - 1; line >= 0; line-- {
		for _, flask := range p.flasks {
			if len(flask.balls) < line+1 {
				// lines = append(lines, string(EMPTY.emoji))
				lines = append(lines, "  ")
			} else {
				lines = append(lines, string(flask.balls[line].emoji))
			}
		}
		lines = append(lines, "\n")
	}

	return strings.Join(lines, "")
}

func (p *Puzzle) String() string {
	if !p.Solved {
		return fmt.Sprintf("Просмотрено состояний - %d. Решений нет\n", len(p.states))
	}

	solutionInText := make([]string, 0, 40)

	p.flasks = p.initialFlasks
	solutionInText = append(solutionInText, "Начальная позиция:", p.getStageStr(), "==============================")

	prev_move, multiplier := p.Moves[0], 1
	for _, move := range p.Moves[1:] {
		if *move == *prev_move {
			multiplier++
		} else {
			if p.confShowMoves {
				p.makeMove(prev_move)
				solutionInText = append(solutionInText, p.getStageStr()[:len(p.getStageStr())-1])
			}
			solutionInText = append(solutionInText, fmt.Sprintf("%dth %s x%d -> %dth tube", prev_move.from+1, prev_move.verboseRuName, multiplier, prev_move.to+1))
			prev_move, multiplier = move, 1
		}
	}
	if p.confShowMoves {
		p.makeMove(prev_move)
		solutionInText = append(solutionInText, p.getStageStr()[:len(p.getStageStr())-1])
	}
	solutionInText = append(solutionInText, fmt.Sprintf("%dth %s x%d -> %dth tube\n", prev_move.from+1, prev_move.verboseRuName, multiplier, prev_move.to+1))

	return fmt.Sprintf("Всего ходов: %d\n%s", len(p.Moves), strings.Join(solutionInText, "\n"))
}
