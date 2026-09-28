package band

func ConsentName(required int) string {
	switch required {
	case 0:
		return "none"
	case 1:
		return "screen sessions"
	}

	return "all sessions"
}
