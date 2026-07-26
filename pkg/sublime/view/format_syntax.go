package view

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/sublime/constant"
)

func (v *View) formatSyntax(_ *option.Format) string {
	if v.Syntax == "" {
		return constant.NoSyntax
	}

	return v.Syntax
}
