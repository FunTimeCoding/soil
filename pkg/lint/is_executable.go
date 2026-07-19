package lint

func IsExecutable(content string) bool {
	if len(content) < 4 {
		return false
	}

	b := []byte(content[:4])

	if b[0] == 0x7f && b[1] == 'E' && b[2] == 'L' && b[3] == 'F' {
		return true
	}

	if b[0] == 0xfe && b[1] == 0xed && b[2] == 0xfa && (b[3] == 0xce || b[3] == 0xcf) {
		return true
	}

	if b[0] == 0xcf && b[1] == 0xfa && b[2] == 0xed && b[3] == 0xfe {
		return true
	}

	if b[0] == 0xce && b[1] == 0xfa && b[2] == 0xed && b[3] == 0xfe {
		return true
	}

	if b[0] == 'M' && b[1] == 'Z' {
		return true
	}

	if b[0] == 0x00 && b[1] == 'a' && b[2] == 's' && b[3] == 'm' {
		return true
	}

	return false
}
