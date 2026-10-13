package band

func RedirectionStateName(state int) string {
	switch state {
	case 32768:
		return "disabled"
	case 32769:
		return "storage"
	case 32770:
		return "serial"
	case 32771:
		return "serial and storage"
	}

	return "unknown"
}
