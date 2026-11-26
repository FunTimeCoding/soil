package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/view"
)

func show(v *view.View) string {
	text := v.Text

	if v.Path != "" {
		text = fmt.Sprintf(
			"snapshot saved to %s (%d lines) - read it in parts",
			v.Path,
			v.Lines,
		)
	}

	if v.Loading {
		return fmt.Sprintf(
			"page still loading after %s\n%s",
			constant.NavigationLimit,
			text,
		)
	}

	return text
}
