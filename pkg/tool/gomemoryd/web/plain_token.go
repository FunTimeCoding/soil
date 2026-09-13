package web

func plainToken(v string) bool {
	if v == "" {
		return false
	}

	for _, r := range v {
		letter := r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z'
		digit := r >= '0' && r <= '9'

		if !letter && !digit && r != '_' && r != '-' {
			return false
		}
	}

	return true
}
