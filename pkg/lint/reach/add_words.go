package reach

import "mvdan.cc/sh/v3/syntax"

func addWords(
	result *CommandWrites,
	directory string,
	words []*syntax.Word,
) {
	for _, w := range words {
		if v, okay := literalValue(w); okay {
			result.Paths = append(result.Paths, absolute(directory, v))

			continue
		}

		if g, okay := globValue(w); okay {
			result.Patterns = append(result.Patterns, absolute(directory, g))

			continue
		}

		result.Unresolved = true
	}
}
