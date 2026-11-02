package loki

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func formatRoute(
	route string,
	f *option.Format,
) string {
	if f.UseColor && route != "" {
		return consoleConstant.Cyan("%s", route)
	}

	return route
}
