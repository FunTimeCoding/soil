package view

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (v *View) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		s.Integer(v.Identifier)
	}

	s.String(v.formatTitle(f), v.formatSyntax(f))

	if v.Dirty {
		s.String(v.formatDirty(f))
	}

	return s.Format()
}
