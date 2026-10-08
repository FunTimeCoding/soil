package mock_random

func New(values ...float64) *Random {
	return &Random{values: values}
}
