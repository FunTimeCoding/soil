package stray_comment

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func isDirective(text string) bool {
	for _, prefix := range []string{
		constant.DirectivePrefix,
		constant.ByValueDirective,
		"//line ",
		"//nolint",
		"//sys",
		"//export",
		"// noinspection ",
		"// renovate: ",
	} {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}

	for _, marker := range []string{"#nosec", "goanalyze:ignore"} {
		if strings.Contains(text, marker) {
			return true
		}
	}

	return false
}
