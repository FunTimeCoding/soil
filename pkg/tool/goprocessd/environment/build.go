package environment

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"strings"
)

func (e *Environment) Build() []string {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	merged := make(map[string]string, len(e.base)+len(e.overlay))

	for _, entry := range e.base {
		key, value, found := strings.Cut(entry, "=")

		if found {
			merged[key] = value
		}
	}

	for key, value := range e.overlay {
		merged[key] = value
	}

	result := make([]string, 0, len(merged))

	for key, value := range merged {
		result = append(result, key_value.Equals(key, value))
	}

	return result
}
