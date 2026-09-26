package service

func priority(v string) bool {
	for _, c := range v {
		if c < '0' || c > '9' {
			return false
		}
	}

	return v != ""
}
