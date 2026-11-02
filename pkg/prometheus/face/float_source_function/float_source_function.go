package float_source_function

import "time"

type FloatSourceFunction func(
	q string,
	t time.Time,
) (float64, error)
