package reach

import "encoding/json"

func inputString(
	raw json.RawMessage,
	key string,
) string {
	var input map[string]any

	if json.Unmarshal(raw, &input) != nil {
		return ""
	}

	result, _ := input[key].(string)

	return result
}
