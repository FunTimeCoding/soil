package overview

func shown(
	show string,
	mitigated bool,
) bool {
	switch show {
	case "all":
		return true
	case "mitigated":
		return mitigated
	default:
		return !mitigated
	}
}
