package guard

import "mvdan.cc/sh/v3/syntax"

type GitCall struct {
	subcommand string
	readable   bool
	arguments  []*syntax.Word
	fromInput  bool
	hidden     bool
}
