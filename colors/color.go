package colors

import (
	"encoding/binary"
	"log"
)

type Color [3]byte

var Red = Color{0xFF, 0x00, 0x00}
var Teal = Color{0x00, 0x80, 0x80}
var Bg = Color{0x00, 0xFF, 0xFF}
var Blue = Color{0x00, 0x00, 0xFF}
var Green = Color{0x00, 0xFF, 0x00}
var Off = Color{0x00, 0x00, 0x00}

func NewColor(hexColor uint32) Color {
	var newColor Color

	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, hexColor)

	newColor[0] = b[2]
	newColor[1] = b[1]
	newColor[2] = b[0]
	log.Println(b[2], b[1], b[0])

	return newColor
}

func SetIntensity(color Color, percentage int) Color {
	var newColor Color

	for i := 0; i < 3; i++ {
		val := (int(color[i]) * percentage) / 100
		newColor[i] = byte(val)

	}

	return newColor
}
