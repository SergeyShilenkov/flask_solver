package puzzle

type Flask struct {
	num   int
	balls [FLASK_SIZE]*Color
}

type UpperBalls struct {
	color  *Color
	amount int
}

func NewFlask(n int, b [FLASK_SIZE]int) *Flask {
	colors := [FLASK_SIZE]*Color{}
	for idx, ball := range b {
		colors[idx] = COLORCONVERT[ball]
	}
	return &Flask{num: n, balls: colors}
}

func (f *Flask) isFull() bool {
	return f.balls[FLASK_SIZE-1] != EMPTY
}

func (f *Flask) isEmpty() bool {
	return f.balls[0] == EMPTY
}

func (f *Flask) freeSpace() int {
	fs := 0
	for i := FLASK_SIZE - 1; i > 0 && f.balls[i] == EMPTY; i-- {
		fs++
	}
	return fs
}

func (f *Flask) hasOneColor() bool {
	result := true
	for _, ball := range f.balls {
		if ball != f.balls[0] && ball != EMPTY {
			result = false
		}
	}
	return result
}

func (f *Flask) isSolved() bool {
	return f.isEmpty() || f.isFull() && f.hasOneColor()
}

func (f *Flask) isAlmostSolved() bool {
	if f.balls[FLASK_SIZE-1] != EMPTY {
		return false
	}

	i := FLASK_SIZE - 2
	for i > 0 && f.balls[i] == f.balls[i-1] {
		i--
	}

	return i == 0
}

func (f *Flask) upperBalls() *UpperBalls {
	i := FLASK_SIZE - 1
	for f.balls[i] == EMPTY {
		i--
	}
	amountWoEmpty := i

	for i > 0 && f.balls[i] == f.balls[i-1] {
		i--
	}
	return &UpperBalls{color: f.balls[amountWoEmpty], amount: amountWoEmpty + 1 - i}
}

func (f *Flask) canReceive(b *Color) bool {
	if f.isFull() {
		return false
	}
	if f.isEmpty() {
		return true
	}
	upper := f.upperBalls()
	return upper.color == b
}

func (f *Flask) pop(n int) *Color {
	var last *Color
	for i := FLASK_SIZE - 1; i >= 0 && n > 0; i-- {
		if f.balls[i] != EMPTY {
			last = f.balls[i]
			f.balls[i] = EMPTY
			n--
		}
	}
	return last
}

func (f *Flask) push(b *Color, n int) {
	for i := 0; i < FLASK_SIZE && n > 0; i++ {
		if f.balls[i] == EMPTY {
			f.balls[i] = b
			n--
		}
	}
}

func (f *Flask) String() []byte {
	state := make([]byte, FLASK_SIZE*len(f.balls[0].Symbol))
	for i := 0; i < FLASK_SIZE; i++ {
		state[2*i] = f.balls[i].Symbol[0]
		state[2*i+1] = f.balls[i].Symbol[1]
	}

	return state
}
