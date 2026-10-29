package face

import "time"

type FloatSource interface {
	QueryFloat(
		q string,
		t time.Time,
	) (float64, error)
}
