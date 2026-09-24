package score

func Score(
	complexity int,
	coverage float64,
) float64 {
	c := float64(complexity)
	missing := 1 - coverage/100

	return c*c*missing*missing*missing + c
}
