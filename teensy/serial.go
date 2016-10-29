package teensy

import (
	"log"

	"Costume/colors"

	"github.com/tarm/serial"
)

const startbyte = '#'

type Serial struct {
	*serial.Port
}

func NewSerial(port string, baud int) *Serial {
	c := &serial.Config{Name: port, Baud: baud}

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	return &Serial{s}

}

func (s *Serial) SetLED(ledNum int, color colors.Color) {
	data := make([]byte, 5)

	data[0] = startbyte
	data[1] = byte(ledNum)
	for i := 0; i < 3; i++ {
		data[i+2] = color[i]
	}
	_, err := s.Write(data)

	log.Printf("%x\n", data)
	if err != nil {
		log.Println("Error writing serial: ", err)
	}
}
