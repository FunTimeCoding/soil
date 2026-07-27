package store

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func withoutSourceType(metadata map[string]string) map[string]string {
	result := map[string]string{}

	for key, value := range metadata {
		if key == constant.SourceType {
			continue
		}

		result[key] = value
	}

	return result
}
