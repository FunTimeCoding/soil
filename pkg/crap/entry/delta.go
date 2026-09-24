package entry

func (e *Entry) Delta() float64 {
	if e.Previous == nil {
		return e.Score
	}

	return e.Score - *e.Previous
}
