package stray_comment

import "strings"

func isDirective(text string) bool {
	for _, prefix := range []string{
		"//go:",
		"//line ",
		"//nolint",
		"//sys",
		"//export",
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
