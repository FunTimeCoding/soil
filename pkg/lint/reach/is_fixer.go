package reach

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"mvdan.cc/sh/v3/syntax"
)

func isFixer(
	name string,
	flags []*syntax.Word,
) bool {
	switch name {
	case "golint", "golangci-lint":
		return hasFlag(flags, constant.FixFlag)
	case "gofix":
		return !hasFlag(flags, constant.DiffFlag)
	case "gofmt", "goimports":
		return hasFlag(flags, constant.WriteFlag)
	case "task":
		return len(flags) > 0 && flags[0].Lit() == "lint"
	}

	return false
}
