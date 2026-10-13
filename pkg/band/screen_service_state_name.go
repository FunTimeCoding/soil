package band

func ScreenServiceStateName(state int) string {
	switch state {
	case 2:
		return "enabled"
	case 3:
		return "disabled"
	case 6:
		return "enabled but offline"
	}

	return "unknown"
}
