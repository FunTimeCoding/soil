package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
)

func answerOf(
	v *client.View,
	e error,
) string {
	if e != nil {
		if c := classify.Class(e); c != "" {
			return c
		}

		return e.Error()
	}

	if v == nil {
		return "ok"
	}

	var text string

	if v.Text != nil {
		text = *v.Text
	}

	return fmt.Sprintf(
		"view %d %q %q %q dirty=%t path=%q",
		v.ViewId,
		v.Title,
		text,
		v.Syntax,
		v.IsDirty,
		v.FilePath,
	)
}
