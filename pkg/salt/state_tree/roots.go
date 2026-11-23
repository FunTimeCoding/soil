package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/types/top_entry"
	"regexp"
	"strings"
)

func roots(entries []*top_entry.Entry) ([]string, []*regexp.Regexp) {
	var result []string
	var patterns []*regexp.Regexp

	for _, e := range entries {
		if strings.Contains(e.State, "{{") {
			patterns = append(patterns, dynamicPattern(e.State))

			continue
		}

		result = append(result, e.State)
	}

	return result, patterns
}
