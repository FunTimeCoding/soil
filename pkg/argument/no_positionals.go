package argument

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
)

func (i *Instance) NoPositionals(hint string) {
	if i.flags.NArg() == 0 {
		return
	}

	system.Exitf(
		1,
		"unexpected arguments: %s\n%s\n",
		join.Space(i.flags.Args()...),
		hint,
	)
}
