package patterns

import (
	"Costume/colors"
	"Costume/teensy"
	"log"
	"time"
)

type Runner struct {
	s            *teensy.Serial
	j            *Jelly
	color        colors.Color
	colorChannel chan colors.Color
	delay        time.Duration
	delayChannel chan time.Duration
	stopChan     chan struct{}
	stoppedChan  chan struct{}
}

func NewRunner(s *teensy.Serial, j *Jelly, color colors.Color, startDelay time.Duration) *Runner {
	return &Runner{
		s:            s,
		j:            j,
		delay:        startDelay,
		colorChannel: make(chan colors.Color),
		delayChannel: make(chan time.Duration),
		stopChan:     make(chan struct{}),
		stoppedChan:  make(chan struct{}),
		color:        color,
	}
}

func (r *Runner) Start() {

	ledsPerStrip := len(r.j.Tenticals[0])
	runner := make([]int, 3)
	runner[0] = 0
	runner[1] = ledsPerStrip - 1
	runner[2] = ledsPerStrip - 2
	currentIntensity := 0

	for i := range r.j.Tenticals {
		for j := range r.j.Tenticals[i] {
			led := r.j.Tenticals[i][j]
			led.Color = colors.Off
			r.s.SetLED(led.Position, led.Color)
		}
	}

	for {
		select {
		case <-r.stopChan:
			return
		//stop
		case newDelay := <-r.delayChannel:
			r.delay = newDelay

		default:

			if currentIntensity == 100 {
				log.Println("HERE")
				for i := range r.j.Tenticals {
					led := r.j.Tenticals[i][runner[2]]
					led.Color = colors.Off
					r.s.SetLED(led.Position, led.Color)
				}

				for i := range runner {

					runner[i] = advanceLed(runner[i], ledsPerStrip)

				}
				log.Println(runner)
				currentIntensity = 0
			}
			currentIntensity += 1

			for i := range r.j.Tenticals {
				led := r.j.Tenticals[i][runner[0]]
				led.Color = colors.SetIntensity(r.color, currentIntensity)
				r.s.SetLED(led.Position, led.Color)

				if runner[0] > 1 {
					led := r.j.Tenticals[i][runner[2]]
					led.Color = colors.SetIntensity(r.color, 100-currentIntensity)
					r.s.SetLED(led.Position, led.Color)
				}
			}

			time.Sleep(r.delay)
		}
	}
}

func (r *Runner) SetColor1(color1 colors.Color) {
	r.colorChannel <- color1
}

//only one color for runner
func (r *Runner) SetColor2(color1 colors.Color) {

}

func (r *Runner) SpeedUp() {
	delay := r.delay
	if delay > (1 * time.Millisecond) {
		r.delayChannel <- delay - 1*time.Millisecond
	}
}
func (r *Runner) SlowDown() {
	delay := r.delay

	r.delayChannel <- delay + 1*time.Millisecond

}

func (r *Runner) Stop() {
	close(r.stopChan)
}

func advanceLed(led int, ledCount int) int {
	led++
	if led == ledCount {
		led = 0

	}
	return led

}
