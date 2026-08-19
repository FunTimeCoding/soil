package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"slices"
)

func matchesMetadata(
	documentMetadata map[string][]string,
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

		if !slices.Contains(documentMetadata[key], value) {
			return false
		}
	}

	return true
}
