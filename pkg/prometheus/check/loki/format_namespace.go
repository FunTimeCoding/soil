package loki

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func formatNamespace(
	e *overview,
	f *option.Format,
) string {
	if f.UseColor && e.Count > 0 {
		return console.Green("%s", e.Namespace)
	}

	return e.Namespace
}
