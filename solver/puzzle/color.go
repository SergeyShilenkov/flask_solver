package puzzle

import (
	"github.com/enescakir/emoji"
)

type Color struct {
	Symbol          string
	verbose_name    string
	verbose_ru_name string
	emoji           emoji.Emoji
}

func (c Color) String() string {
	return string(c.emoji)
}

var (
	RED    = &Color{"R", "Red", "Красный", emoji.RedCircle}
	GREEN  = &Color{"G", "Green", "Зелёный", emoji.GreenCircle}
	ORANGE = &Color{"O", "Orange", "Оранжевый", emoji.OrangeCircle}
	BLUE   = &Color{"B", "Blue", "Синий", emoji.BlueCircle}
	PURPLE = &Color{"V", "Purple", "Фиолетовый", emoji.PurpleCircle}
	YELLOW = &Color{"Y", "Yellow", "Жёлтый", emoji.YellowCircle}
	GRAY   = &Color{"Q", "Gray", "Серый", emoji.WhiteCircle}
	BROWN  = &Color{"X", "Brown", "Коричневый", emoji.BrownCircle}

	PINK      = &Color{"P", "Pink", "Розовый", emoji.Brain}
	LILAC     = &Color{"Pu", "Lilac", "Сиреневый", emoji.PurpleHeart}
	CRIMSON   = &Color{"c", "Crimson", "Малиновый", emoji.RedHeart}
	TURQUOISE = &Color{"b", "Turquoise", "Бирюзовый", emoji.BlueHeart}
	L_GREEN   = &Color{"g", "Light Green", "Светло-зелёный", emoji.GreenHeart}

	D_BLUE = &Color{"BB", "Dark blue", "Тёмно-синий", emoji.BlueSquare}

	UNKNOWN = &Color{"U", "Unknown", "Неизвестный", emoji.QuestionMark}

	EMPTY = &Color{"E", "Empty", "-", emoji.TumblerGlass}

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
