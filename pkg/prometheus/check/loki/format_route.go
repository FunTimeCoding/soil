package loki

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func formatRoute(
	route string,
	f *option.Format,
) string {
	if f.UseColor && route != "" {
		return console.Cyan("%s", route)
	}

	return route
}
