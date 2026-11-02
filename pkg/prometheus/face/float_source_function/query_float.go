package float_source_function

import "time"

func (f FloatSourceFunction) QueryFloat(
	q string,
	t time.Time,
) (float64, error) {
	return f(q, t)
}
