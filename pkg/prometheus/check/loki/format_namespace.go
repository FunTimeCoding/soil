package loki

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func formatNamespace(
	e *overview,
	f *option.Format,
) string {
	if f.UseColor && e.Count > 0 {
		return constant.Green("%s", e.Namespace)
	}

	return e.Namespace
}
