package golinkace

import "strings"

func splitTrimmed(v string) []string {
	parts := strings.Split(v, ",")
	var result []string

	for _, p := range parts {
		t := strings.TrimSpace(p)

		if t != "" {
			result = append(result, t)
		}
	}

	return result
}
