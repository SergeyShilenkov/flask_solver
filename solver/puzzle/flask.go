package puzzle

import "strings"

type Flask struct {
	num   int
	balls []*Color
}

type UpperBalls struct {
	color  *Color
	amount int
}

func NewFlask(n int, b []int) *Flask {
	colors := make([]*Color, 0, FLASK_SIZE)
	for _, ball := range b {
		if ball != 0 {
			colors = append(colors, COLORCONVERT[ball])
		}
	}
	return &Flask{num: n, balls: colors}
}

func (f *Flask) isFull() bool {
	return len(f.balls) == FLASK_SIZE
}

func (f *Flask) isEmpty() bool {
	return len(f.balls) == 0
}

func (f *Flask) freeSpace() int {
	return FLASK_SIZE - len(f.balls)
}

func (f *Flask) hasOneColor() bool {
	result := true
	for _, ball := range f.balls {
		if *ball != *(f.balls[0]) {
			result = false
		}
	}
	return result
}

func (f *Flask) isSolved() bool {
	return f.isEmpty() || f.isFull() && f.hasOneColor()
}

func (f *Flask) isAlmostSolved() bool {
	return len(f.balls) == FLASK_SIZE-1 && f.hasOneColor()
}

func (f *Flask) upperBalls() *UpperBalls {
	i := len(f.balls) - 1
	for i > 0 && f.balls[i] == f.balls[i-1] {
		i--
	}
	return &UpperBalls{color: f.balls[len(f.balls)-1], amount: len(f.balls) - i}
}

func (f *Flask) canReceive(b *Color) bool {
	if f.isFull() {
		return false
	}
	if f.isEmpty() {
		return true
	}
	upper := f.upperBalls()
	return *(upper.color) == *b
}

func (f *Flask) pop() *Color {
	last := f.balls[len(f.balls)-1]
	f.balls = f.balls[:len(f.balls)-1]
	return last
}

func (f *Flask) push(b *Color) {
	f.balls = append(f.balls, b)
}

func (f *Flask) String() string {
	state := make([]string, FLASK_SIZE)
	for i := 0; i < FLASK_SIZE; i++ {
		if i <= len(f.balls)-1 {
			state[i] = f.balls[i].Symbol
		} else {
			state[i] = "-"
		}
	}
	return strings.Join(state, "")
}
