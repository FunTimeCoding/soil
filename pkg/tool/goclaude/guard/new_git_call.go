package guard

import "mvdan.cc/sh/v3/syntax"

func newGitCall(
	subcommand string,
	readable bool,
	arguments []*syntax.Word,
	fromInput bool,
) *GitCall {
	return &GitCall{
		subcommand: subcommand,
		readable:   readable,
		arguments:  arguments,
		fromInput:  fromInput,
	}
}
