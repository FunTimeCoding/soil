package loader

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/maps"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (l *Loader) ToText(m map[string]map[string]string) string {
	var result []string

	for k, v := range m {
		result = append(result, IntegerFromString(k))

		for _, kk := range maps.StringKeys(v) {
			result = append(result, fmt.Sprintf("%s: %s", kk, v[kk]))
		}
	}

	return join.NewLine(result)
}
