package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"path"
	"slices"
	"strings"
)

func findGitCall(call *syntax.CallExpr) *GitCall {
	words := call.Args
	fromInput := false

	if path.Base(words[0].Lit()) == "xargs" {
		fromInput = true
		words = words[1:]

		for len(words) > 0 && strings.HasPrefix(words[0].Lit(), "-") {
			words = words[1:]
		}
	}

	if len(words) == 0 || path.Base(words[0].Lit()) != "git" {
		return nil
	}

	words = words[1:]

	for len(words) > 0 {
		option, literal := literalWord(words[0])

		if !literal || !strings.HasPrefix(option, "-") {
			break
		}

		words = words[1:]

		if slices.Contains(constant.GitOptionsWithValue, option) &&
			len(words) > 0 {
			words = words[1:]
		}
	}

	if len(words) == 0 {
		return nil
	}

	subcommand, literal := literalWord(words[0])

	return newGitCall(subcommand, literal, words[1:], fromInput)
}
