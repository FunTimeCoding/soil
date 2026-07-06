package view

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (v *View) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer(v.Identifier)
	}

	s.String(v.formatTitle(f), v.formatSyntax(f))

	if v.Dirty {
		s.String(v.formatDirty(f))
	}

	return s.Format()
}
