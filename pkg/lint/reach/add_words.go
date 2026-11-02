package reach

import "mvdan.cc/sh/v3/syntax"

func addWords(
	result *CommandWrites,
	directory string,
	words []*syntax.Word,
) {
	for _, w := range words {
		v, okay := literalValue(w)

		if !okay {
			result.Unresolved = true

			continue
		}

		result.Paths = append(result.Paths, absolute(directory, v))
	}
}
