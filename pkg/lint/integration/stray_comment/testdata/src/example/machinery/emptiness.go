package machinery

type Marker interface {
	// marker
}

func Dispatch(kind int) string {
	switch kind {
	case 0:
		// pass
	case 1:
		return "one"
	default:
		// pass
	}

	if kind > 1 {
		// pass
	}

	// pass the token along
	return ""
}
