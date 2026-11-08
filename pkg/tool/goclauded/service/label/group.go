package label

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
)

func group(rows []label.Label) []label.Label {
	var result []label.Label
	var values []string

	for i, r := range rows {
		values = append(values, r.Value)

		if i+1 < len(rows) && rows[i+1].Key == r.Key {
			continue
		}

		result = append(
			result,
			*label.New(r.SessionIdentifier, r.Key, join.Comma(values)),
		)
		values = nil
	}

	return result
}
