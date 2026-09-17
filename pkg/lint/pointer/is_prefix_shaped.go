package pointer

import "unicode"

func isPrefixShaped(s string) bool {
	digits := 0

	for i, r := range s {
		switch {
		case unicode.IsLower(r) && digits == 0:
			continue
		case unicode.IsDigit(r) && i > 0:
			digits++
		default:
			return false
		}
	}

	return digits > 0
}
