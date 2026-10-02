package goflow

import (
	"github.com/funtimecoding/soil/pkg/tool/goflow/constant"
	"strings"
)

func units(
	source []byte,
	literal []bool,
	start int,
	stop int,
) []string {
	var result []string
	var unit []byte

	for i := start; i < stop; i++ {
		c := source[i]

		if !literal[i] && strings.IndexByte(constant.Whitespace, c) >= 0 {
			if len(unit) > 0 {
				result = append(result, string(unit))
				unit = nil
			}

			continue
		}

		if literal[i] && c == '\n' {
			c = ' '

			for i+1 < stop && strings.IndexByte(constant.Indent, source[i+1]) >= 0 {
				i++
			}
		}

		unit = append(unit, c)
	}

	if len(unit) > 0 {
		result = append(result, string(unit))
	}

	return result
}
