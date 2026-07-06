package sed

import "github.com/funtimecoding/soil/pkg/strings/split/key_value"

func Parse(replaces []string) map[string]string {
	result := make(map[string]string)

	for _, r := range replaces {
		k, v := key_value.Equals(r)
		result[k] = v
	}

	return result
}
