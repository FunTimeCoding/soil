package mock_client

import "strings"

func containsType(
	types []string,
	t string,
) bool {
	for _, v := range types {
		if strings.EqualFold(v, t) {
			return true
		}
	}

	return false
}
