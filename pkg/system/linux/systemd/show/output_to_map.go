package show

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
)

func OutputToMap(s string) map[string]string {
	result := make(map[string]string)

	for _, l := range split.NewLine(s) {
		k, v := key_value.Equals(l)
		result[k] = v
	}

	return result
}
