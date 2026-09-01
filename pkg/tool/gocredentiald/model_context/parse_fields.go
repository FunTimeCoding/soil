package model_context

import "strings"

func parseFields(pairs []string) (map[string]string, string) {
	result := map[string]string{}

	for _, pair := range pairs {
		key, value, found := strings.Cut(pair, "=")

		if !found || key == "" {
			return nil, pair
		}

		result[key] = value
	}

	return result, ""
}
