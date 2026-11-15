package shell_script

import (
	"mvdan.cc/sh/v3/syntax"
	"strings"
)

func nestedScript(
	words []*syntax.Word,
	parent *Script,
) *Script {
	printed := make([]string, 0, len(words))
	readable := true

	for _, w := range words {
		printed = append(printed, Printed(w))

		if _, okay := Text(w); !okay {
			readable = false
		}
	}

	return &Script{
		Text:     scriptText(words),
		Printed:  strings.Join(printed, " "),
		Readable: readable,
		Parent:   parent,
	}
}
