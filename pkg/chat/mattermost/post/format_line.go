package post

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func formatLine(
	f *option.Format,
	result string,
) string {
	if result == "" {
		result = "(no text, image-only)"
	}

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
