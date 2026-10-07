package lint

import "github.com/funtimecoding/soil/pkg/lint/option"

func Skipped(
	o *option.Lint,
	path string,
) bool {
	return SkippedBy(o.Skips, path)
}
