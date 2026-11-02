package loki

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func formatNamespace(
	e *Overview,
	f *option.Format,
) string {
	if f.UseColor && e.Count > 0 {
		return consoleConstant.Green("%s", e.Namespace)
	}

	return e.Namespace
}
