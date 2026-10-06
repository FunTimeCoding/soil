package requester

import "github.com/funtimecoding/soil/pkg/web/constant"

func text(v any) string {
	switch t := v.(type) {
	case string:
		return t
	case []any:
		for _, item := range t {
			if s := text(item); s != "" {
				return s
			}
		}
	case map[string]any:
		for _, k := range constant.NestedFields {
			if s := text(t[k]); s != "" {
				return s
			}
		}
	}

	return ""
}
