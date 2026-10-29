package refusal

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"sort"
)

func reason(
	body []byte,
	fallback string,
) string {
	var detail Detail

	if json.Unmarshal(body, &detail) == nil && detail.Detail != "" {
		return detail.Detail
	}

	var fields map[string][]string

	if json.Unmarshal(body, &fields) != nil || len(fields) == 0 {
		return fallback
	}

	var parts []string

	for field, messages := range fields {
		for _, message := range messages {
			if field == "__all__" {
				parts = append(parts, message)
			} else {
				parts = append(parts, fmt.Sprintf("%s: %s", field, message))
			}
		}
	}

	sort.Strings(parts)

	return join.SemicolonSpace(parts)
}
