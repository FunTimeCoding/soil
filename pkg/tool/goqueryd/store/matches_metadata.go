package store

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func matchesMetadata(
	documentMetadata map[string]string,
	resolvedSourceType string,
	filter map[string]string,
) bool {
	for key, value := range filter {
		if key == constant.SourceType {
			if resolvedSourceType != value {
				return false
			}

			continue
		}

		if documentMetadata == nil || documentMetadata[key] != value {
			return false
		}
	}

	return true
}
