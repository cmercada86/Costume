package patterns

import "Costume/colors"
import "log"

type LED struct {
	Position int
	Color    colors.Color
}

type Jelly struct {
	Tenticals [][]LED
	Rim       []LED
}

func NewJelly(numTenticals int, ledsPerStrip int, ledsOnRim int) *Jelly {

	tenticals := make([][]LED, numTenticals)
	for i := range tenticals {
		tenticals[i] = make([]LED, ledsPerStrip)
		for j := range tenticals[i] {
			tenticals[i][j] = LED{
				Position: i*ledsPerStrip + j,
				Color:    colors.Off,
			}
			log.Println(tenticals[i][j])
		}
	}
	rim := make([]LED, ledsOnRim)
	for i := range rim {
		rim[i] = LED{
			Position: numTenticals*ledsPerStrip + i,
			Color:    colors.Off,
		}
	}

	return &Jelly{
		Tenticals: tenticals,
		Rim:       rim,
	}
}
