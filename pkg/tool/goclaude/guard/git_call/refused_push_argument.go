package git_call

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"slices"
	"strings"
)

func refusedPushArgument(argument string) bool {
	if strings.HasPrefix(argument, "+") || strings.HasPrefix(argument, ":") {
		return true
	}

	if strings.HasPrefix(argument, "--") {
		name, _, _ := strings.Cut(argument, "=")

		return slices.Contains(constant.GitRefusedPushOptions, name)
	}

	if strings.HasPrefix(argument, "-") {
		return strings.ContainsAny(argument[1:], constant.GitRefusedPushShort)
	}

	return false
}
