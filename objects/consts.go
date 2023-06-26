package objects

import "image/color"

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
	SQUARE_SIZE   = 10

	LEFT_BORDER   = 40
	RIGHT_BORDER  = 760
	BOTTOM_BORDER = 560
	UPPER_BORDER  = 160
)

var (
	RED   = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	GREEN = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	WHITE = color.White
)
