package state_tree

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"regexp"
	"strings"
)

func referenceOf(path string) (string, *regexp.Regexp) {
	path = strings.TrimSuffix(path, "/")

	if strings.Contains(path, "{{") {
		return "", dynamicPattern(join.Empty(path, "{{}}"))
	}

	return path, nil
}
