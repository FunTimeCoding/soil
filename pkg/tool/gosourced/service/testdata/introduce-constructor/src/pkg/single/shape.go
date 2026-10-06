package single

import "time"

type Shape struct {
	Draw   float64
	Inhale float64
	Hold   float64
	Wait   time.Duration
}
