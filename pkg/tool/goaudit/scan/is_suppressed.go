package scan

func isSuppressed(
	suppress map[string][]string,
	operationIdentifier string,
	concernKey string,
) bool {
	keys, okay := suppress[operationIdentifier]

	if !okay {
		return false
	}

	for _, k := range keys {
		if k == concernKey {
			return true
		}
	}

	return false
}
