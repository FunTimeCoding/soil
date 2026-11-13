package component

import "github.com/funtimecoding/soil/pkg/web/constant"

func OrUnset(value string) string {
	if value == "" {
		return constant.UnsetMark
	}

	return value
}
