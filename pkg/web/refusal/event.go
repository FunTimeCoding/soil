package refusal

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func Event(body []byte) string {
	var fields map[string]any

	if json.Unmarshal(body, &fields) != nil {
		return ""
	}

	result, _ := fields[constant.EventField].(string)

	return result
}
