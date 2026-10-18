package refusal

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/web/constant"
)

func Reason(
	body []byte,
	status string,
) string {
	var fields map[string]any

	if json.Unmarshal(body, &fields) != nil {
		return status
	}

	for _, k := range constant.ReasonFields {
		if v := text(fields[k]); v != "" {
			return v
		}
	}

	return status
}
