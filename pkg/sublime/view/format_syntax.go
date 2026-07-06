package view

import "github.com/funtimecoding/soil/pkg/console/status/option"

func (v *View) formatSyntax(_ *option.Format) string {
	if v.Syntax == "" {
		return NoSyntax
	}

	return v.Syntax
}
