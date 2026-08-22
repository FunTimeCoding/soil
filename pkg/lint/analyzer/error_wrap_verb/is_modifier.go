package error_wrap_verb

func isModifier(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}

	switch c {
	case '+', '-', '#', ' ', '.', '*':
		return true
	}

	return false
}
