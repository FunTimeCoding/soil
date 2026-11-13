package outage_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"slices"
)

func declared[V any](
	served []string,
	exceptions map[string]constant.OutageReach,
	overrides map[string]V,
) []string {
	var result []string

	for name := range exceptions {
		if !slices.Contains(served, name) {
			result = append(
				result,
				fmt.Sprintf("%s: declared, not served", name),
			)
		}
	}

	for name := range overrides {
		if !slices.Contains(served, name) {
			result = append(
				result,
				fmt.Sprintf("%s: override, not served", name),
			)
		}
	}

	return result
}
