package installed

func Moved(
	built map[string]string,
	current map[string]string,
) bool {
	for module, version := range built {
		if now, okay := current[module]; okay && now != version {
			return true
		}
	}

	return false
}
