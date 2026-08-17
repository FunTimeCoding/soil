package service

import "encoding/json"

func inputString(
	raw string,
	key string,
) string {
	var input map[string]any

	if json.Unmarshal([]byte(raw), &input) != nil {
		return ""
	}

	value, okay := input[key].(string)

	if !okay {
		return ""
	}

	return value
}
