package environment

func SplitNull(
	content []byte,
	atEOF bool,
) (int, []byte, error) {
	for i, b := range content {
		if b == 0 {
			return i + 1, content[:i], nil
		}
	}

	if atEOF && len(content) > 0 {
		return len(content), content, nil
	}

	return 0, nil, nil
}
