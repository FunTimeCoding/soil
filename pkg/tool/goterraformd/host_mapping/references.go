package host_mapping

import "github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"

func references(v any) []string {
	var result []string

	switch t := v.(type) {
	case map[string]any:
		for key, value := range t {
			if key != constant.ReferenceKey {
				result = append(result, references(value)...)

				continue
			}

			if list, okay := value.([]any); okay {
				for _, r := range list {
					if s, isString := r.(string); isString {
						result = append(result, s)
					}
				}
			}
		}
	case []any:
		for _, e := range t {
			result = append(result, references(e)...)
		}
	}

	return result
}
