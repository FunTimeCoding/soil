package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
	"path/filepath"
	"slices"
)

func restores(
	call *syntax.CallExpr,
	target string,
) bool {
	var operands []string

	for _, a := range call.Args[1:] {
		if value, okay := wordValue(a); okay {
			operands = append(operands, filepath.Clean(value))
		}
	}

	if len(operands) == 0 {
		return false
	}

	switch path.Base(call.Args[0].Lit()) {
	case "cp", "mv":
		return operands[len(operands)-1] == target
	case "git":
		return (operands[0] == "checkout" || operands[0] == "restore") &&
			slices.Contains(operands[1:], target)
	default:
		return false
	}
}
