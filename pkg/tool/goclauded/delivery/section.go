package delivery

func section(
	header string,
	lines []string,
) []string {
	if len(lines) == 0 {
		return nil
	}

	return append([]string{header}, lines...)
}
