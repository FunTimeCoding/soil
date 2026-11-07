package guard

import (
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"mvdan.cc/sh/v3/syntax"
	"path"
	"slices"
	"strings"
)

func wrappedCommand(words []*syntax.Word) ([]*syntax.Word, bool) {
	fromInput := false

	for len(words) > 0 {
		name, literal := literalWord(words[0])
		name = path.Base(name)

		if !literal || !slices.Contains(constant.CommandWrappers, name) {
			return words, fromInput
		}

		if name == constant.XargsCommand {
			fromInput = true
		}

		words = words[1:]

		for len(words) > 0 {
			option, _ := literalWord(words[0])

			if !strings.HasPrefix(option, "-") {
				break
			}

			words = words[1:]

			if slices.Contains(constant.WrapperValueOptions[name], option) &&
				len(words) > 0 {
				words = words[1:]
			}
		}

		if name == constant.TimeoutCommand && len(words) > 0 {
			words = words[1:]
		}
	}

	return words, fromInput
}
