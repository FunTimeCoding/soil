package classify

func Reportable(e error) bool {
	if a := Answer(e); a != nil {
		return a.Captured
	}

	return !expected(e)
}
