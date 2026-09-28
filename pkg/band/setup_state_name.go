package band

func SetupStateName(state int) string {
	switch state {
	case 0:
		return "pre-provisioning"
	case 1:
		return "in-provisioning"
	case 2:
		return "provisioned"
	}

	return "unknown"
}
