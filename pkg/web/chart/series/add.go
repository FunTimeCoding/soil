package series

import "time"

func (s *Series) Add(
	at time.Time,
	value float64,
) *Series {
	s.Times = append(s.Times, at)
	s.Value = append(s.Value, value)

	return s
}
