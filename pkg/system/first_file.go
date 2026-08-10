package system

func FirstFile(s ...string) string {
	for _, path := range s {
		if FileExists(path) {
			return path
		}
	}

	return ""
}
