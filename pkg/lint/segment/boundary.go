package segment

import "unicode"

func boundary(
	r []rune,
	i int,
) bool {
	previous := r[i-1]
	current := r[i]

	if !unicode.IsUpper(previous) && unicode.IsUpper(current) {
		return true
	}

	if !unicode.IsUpper(previous) || !unicode.IsUpper(current) ||
		i+1 >= len(r) || !unicode.IsLower(r[i+1]) {
		return false
	}

	if r[i+1] == 's' && (i+2 == len(r) || unicode.IsUpper(r[i+2])) {
		return false
	}

	return true
}
