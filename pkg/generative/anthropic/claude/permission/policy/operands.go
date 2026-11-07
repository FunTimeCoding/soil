package policy

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func operands(arguments []*syntax.Word) ([]*syntax.Word, []string) {
	var result []*syntax.Word
	var options []string
	ended := false

	for _, a := range arguments {
		value, okay := literal(a)

		if !ended && okay && value == constant.OptionTerminator {
			ended = true

			continue
		}

		if !ended && okay && strings.HasPrefix(value, constant.OptionPrefix) &&
			value != constant.OptionPrefix {
			options = append(options, value)

			continue
		}

		result = append(result, a)
	}

	return result, options
}
