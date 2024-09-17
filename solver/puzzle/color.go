package puzzle

import (
	"github.com/enescakir/emoji"
)

type Color struct {
	Symbol  [2]byte
	verbose string
	emoji   emoji.Emoji
}

func (c Color) String() string {
	return string(c.emoji)
}

var (
	RED    = &Color{[2]byte{'-', 'R'}, "Красный", emoji.RedCircle}
	GREEN  = &Color{[2]byte{'-', 'G'}, "Зелёный", emoji.GreenCircle}
	ORANGE = &Color{[2]byte{'-', 'O'}, "Оранжевый", emoji.OrangeCircle}
	BLUE   = &Color{[2]byte{'-', 'B'}, "Синий", emoji.BlueCircle}
	PURPLE = &Color{[2]byte{'-', 'V'}, "Фиолетовый", emoji.PurpleCircle}
	YELLOW = &Color{[2]byte{'-', 'Y'}, "Жёлтый", emoji.YellowCircle}
	GRAY   = &Color{[2]byte{'-', 'Q'}, "Серый", emoji.WhiteCircle}
	BROWN  = &Color{[2]byte{'-', 'X'}, "Коричневый", emoji.BrownCircle}

	PINK      = &Color{[2]byte{'-', 'P'}, "Розовый", emoji.Brain}
	LILAC     = &Color{[2]byte{'P', 'u'}, "Сиреневый", emoji.PurpleHeart}
	CRIMSON   = &Color{[2]byte{'-', 'c'}, "Малиновый", emoji.RedHeart}
	TURQUOISE = &Color{[2]byte{'-', 'b'}, "Бирюзовый", emoji.BlueHeart}
	L_GREEN   = &Color{[2]byte{'-', 'g'}, "Светло-зелёный", emoji.GreenHeart}

	D_BLUE = &Color{[2]byte{'B', 'B'}, "Тёмно-синий", emoji.BlueSquare}

	UNKNOWN = &Color{[2]byte{'-', 'U'}, "Неизвестный", emoji.QuestionMark}

	EMPTY = &Color{[2]byte{'-', 'E'}, "-", emoji.TumblerGlass}

	COLORCONVERT = map[int]*Color{
		0:  EMPTY,
		1:  UNKNOWN,
		2:  TURQUOISE,
		3:  YELLOW,
		4:  GREEN,
		5:  BROWN,
		6:  RED,
		7:  CRIMSON,
		8:  ORANGE,
		9:  PINK,
		10: L_GREEN,
		11: GRAY,
		12: BLUE,
		13: LILAC,
		14: D_BLUE,
		15: PURPLE,
	}
)
