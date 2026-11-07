package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"slices"
	"strings"
)

func escapes(
	name string,
	arguments []*syntax.Word,
) bool {
	if slices.Contains(constant.GlobSensitiveCommands, name) &&
		unquotedGlob(arguments) {
		return true
	}

	if name != constant.FindCommand {
		return false
	}

	for _, a := range arguments {
		value, okay := literal(a)

		if !okay {
			return true
		}

		for _, o := range constant.FindEscapes {
			if strings.HasPrefix(value, o) {
				return true
			}
		}
	}

	return false
}
