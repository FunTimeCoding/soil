package gocheck

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func matchesPattern(
	relative string,
	pattern string,
) bool {
	if pattern == goModConstant.AllPackages {
		return true
	}

	trimmed := strings.TrimPrefix(
		pattern,
		join.Empty(soil.CurrentDirectory, stringsConstant.Slash),
	)
	base, recursive := strings.CutSuffix(
		trimmed,
		goModConstant.RecursivePattern,
	)

	if !recursive {
		return relative == trimmed
	}

	return relative == base ||
		strings.HasPrefix(relative, join.Empty(base, stringsConstant.Slash))
}
