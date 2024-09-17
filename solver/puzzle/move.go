package puzzle

import (
	"fmt"

	"github.com/enescakir/emoji"
)

type Move struct {
	from       int
	to         int
	emoji      emoji.Emoji
	verbose    string
	ballAmount int
}

func (m Move) String() string {
	return fmt.Sprintf("%d -> %d x%d", m.from, m.to, m.ballAmount)
}

func NewMove(f int, t int, u *UpperBalls) *Move {
	return &Move{from: f, to: t, emoji: u.color.emoji, verbose: u.color.verbose, ballAmount: u.amount}
}
