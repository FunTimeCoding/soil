package reach

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/lint/reach/transcript"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func resultPaths(raw json.RawMessage) []string {
	var texts []string
	var single string

	if json.Unmarshal(raw, &single) == nil {
		texts = append(texts, single)
	} else {
		var parts []*transcript.Text

		if json.Unmarshal(raw, &parts) == nil {
			for _, p := range parts {
				texts = append(texts, p.Text)
			}
		}
	}

	var result []string

	for _, text := range texts {
		for _, line := range split.NewLine(text) {
			result = append(result, pathsIn(line)...)
		}
	}

	return result
}
