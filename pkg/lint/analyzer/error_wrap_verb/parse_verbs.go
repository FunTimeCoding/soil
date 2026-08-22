package error_wrap_verb

func parseVerbs(format string) (map[int]byte, bool) {
	verbs := map[int]byte{}
	argument := 0

	for i := 0; i < len(format); i++ {
		if format[i] != '%' {
			continue
		}

		i++

		if i >= len(format) {
			break
		}

		if format[i] == '%' {
			continue
		}

		for i < len(format) {
			if format[i] == '[' {
				return nil, false
			}

			if !isModifier(format[i]) {
				break
			}

			if format[i] == '*' {
				argument++
			}

			i++
		}

		if i >= len(format) {
			break
		}

		verbs[argument] = format[i]
		argument++
	}

	return verbs, true
}
