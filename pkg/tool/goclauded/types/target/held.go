package target

import "strings"

func (t *Target) Held(
	key string,
	value string,
) string {
	for _, v := range t.Labels[key] {
		if strings.EqualFold(v, value) {
			return v
		}
	}

	return ""
}
