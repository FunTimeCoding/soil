package guard

import (
	"mvdan.cc/sh/v3/syntax"
	"path"
)

func pipInstallCall(call *syntax.CallExpr) bool {
	base := path.Base(call.Args[0].Lit())
	pip := base == "pip" || base == "pip3"

	if !pip && base != "python" && base != "python3" {
		return false
	}

	module := false

	for i, argument := range call.Args[1:] {
		literal := argument.Lit()

		if !pip && literal == "-m" && i+2 < len(call.Args) {
			next := call.Args[i+2].Lit()

			if next == "pip" || next == "pip3" {
				module = true
			}
		}

		if literal == "install" && (pip || module) {
			return true
		}
	}

	return false
}
