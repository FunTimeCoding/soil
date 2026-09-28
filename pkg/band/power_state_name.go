package band

func PowerStateName(state int) string {
	switch state {
	case 2:
		return "on"
	case 3, 4:
		return "sleep"
	case 5:
		return "power cycle"
	case 6:
		return "off hard"
	case 7:
		return "hibernate"
	case 8:
		return "off soft"
	case 10:
		return "reset"
	case 11:
		return "diagnostic interrupt"
	case 12:
		return "off soft graceful"
	case 13:
		return "off hard graceful"
	case 14:
		return "reset graceful"
	case 15:
		return "power cycle graceful"
	}

	return "unknown"
}
