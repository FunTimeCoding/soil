package git_call

import "mvdan.cc/sh/v3/syntax"

func New(
	subcommand string,
	readable bool,
	arguments []*syntax.Word,
	fromInput bool,
) *Call {
	return &Call{
		subcommand: subcommand,
		readable:   readable,
		arguments:  arguments,
		fromInput:  fromInput,
	}
}
