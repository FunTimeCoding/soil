package board

func Port(t Target) int {
	if t.Port != 0 {
		return t.Port
	}

	if t.Secure {
		return securePort
	}

	return insecurePort
}
