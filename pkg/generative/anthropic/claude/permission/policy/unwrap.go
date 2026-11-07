package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"slices"
	"strings"
)

func (w *Walk) unwrap(
	name string,
	arguments []*syntax.Word,
) bool {
	i := 0

	for i < len(arguments) {
		value, okay := literal(arguments[i])

		if !okay {
			return w.block(name)
		}

		if !strings.HasPrefix(value, constant.OptionPrefix) {
			break
		}

		if name == constant.XargsCommand || name == constant.CommandBuiltin {
			return w.block(
				w.commandLabel(
					append([]string{name}, literalWords(arguments)...),
				),
			)
		}

		i++

		if slices.Contains(constant.WrapperValueOptions[name], value) {
			i++
		}
	}

	if name == constant.TimeoutCommand {
		i++
	}

	if i >= len(arguments) {
		return w.block(name)
	}

	return w.invoke(arguments[i:])
}
