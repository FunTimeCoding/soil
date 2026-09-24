package crap

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/crap/constant"
	"github.com/funtimecoding/soil/pkg/crap/entry"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func label(e *entry.Entry) string {
	name := e.Function.QualifiedName()

	if e.Missing {
		name = join.Empty(name, constant.MissingSuffix)
	}

	if e.Untrusted {
		name = join.Empty(
			name,
			fmt.Sprintf(
				constant.UntrustedFormat,
				e.Killed,
				e.Killed+len(e.Lived),
			),
		)
	}

	return name
}
