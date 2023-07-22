package objects

import (
	"golang.org/x/image/colornames"
)

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
	SQUARE_SIZE   = 20

	LEFT_BORDER   = 40
	RIGHT_BORDER  = SCREEN_WIDTH - LEFT_BORDER - SQUARE_SIZE
	BOTTOM_BORDER = 560
	UPPER_BORDER  = 160
)

var (
	RED   = colornames.Red
	GREEN = colornames.Green
	WHITE = colornames.White
	BLACK = colornames.Black
)
