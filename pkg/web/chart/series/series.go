package series

import "time"

type Series struct {
	Label string
	Class string
	Times []time.Time
	Value []float64
}
