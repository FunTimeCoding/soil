package search_index

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func inputText(
	name string,
	input json.RawMessage,
) string {
	var v any

	if json.Unmarshal(input, &v) != nil {
		return name
	}

	return join.NewLine(append([]string{name}, stringValues(v)...))
}
