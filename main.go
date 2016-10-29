package main

import (
	"Costume/colors"
	"Costume/patterns"
	"Costume/teensy"
	"time"
)

func main() {
	s := teensy.NewSerial("COM3", 115200)

	jelly := patterns.NewJelly(6, 9, 0)

	runner := patterns.NewRunner(s, jelly, colors.Bg, 5*time.Millisecond)

	runner.Start()

}
