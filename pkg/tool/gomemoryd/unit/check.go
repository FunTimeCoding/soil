package unit

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/reference"

func check(content string) []*reference.Finding {
	return checkUnder(content)
}
