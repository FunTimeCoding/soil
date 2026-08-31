package service

import "github.com/funtimecoding/soil/pkg/tool/gosourced/service/match"

func parseApplyPatterns(
	pattern string,
	replacement string,
	symbol string,
) (*match.Pattern, *match.Pattern, error) {
	patternSpec, e := match.Parse(pattern)

	if e != nil {
		return nil, nil, e
	}

	replacementSpec, f := match.Parse(replacement)

	if f != nil {
		return nil, nil, f
	}

	if g := match.CheckApplyShape(patternSpec, replacementSpec, symbol); g != nil {
		return nil, nil, g
	}

	return patternSpec, replacementSpec, nil
}
