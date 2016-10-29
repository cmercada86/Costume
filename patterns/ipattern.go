package patterns

import "Costume/colors"

type Ipattern interface {
	Start()
	Stop()
	SetColor1(color1 colors.Color)
	SetColor2(color2 colors.Color)
	SpeedUp()
	SlowDown()
}
