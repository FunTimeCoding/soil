package meter

func New(
	value float64,
	maximum float64,
) *Meter {
	return &Meter{value: value, maximum: maximum}
}
