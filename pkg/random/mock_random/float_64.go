package mock_random

func (r *Random) Float64() float64 {
	value := r.values[r.index%len(r.values)]
	r.index++

	return value
}
