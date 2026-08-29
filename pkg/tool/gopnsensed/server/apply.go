package server

func apply(v *bool) bool {
	if v == nil {
		return true
	}

	return *v
}
