package goclaude

import "github.com/funtimecoding/soil/pkg/tool/goclaude/constant"

func shortModelName(display string) string {
	if short, okay := constant.ShortModelName[display]; okay {
		return short
	}

	return display
}
