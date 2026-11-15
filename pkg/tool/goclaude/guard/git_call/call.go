package git_call

import "mvdan.cc/sh/v3/syntax"

type Call struct {
	subcommand string
	readable   bool
	arguments  []*syntax.Word
	fromInput  bool
	hidden     bool
}
