package lint

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
)

func scopeDetail(
	o *option.Lint,
	v *virtual_file_system.System,
) string {
	if len(o.Scopes) == 0 {
		return ""
	}

	count := 0

	for _, p := range v.Files() {
		if InScope(o, p) && !Skipped(o, p) {
			count++
		}
	}

	return fmt.Sprintf("scope %s (%d selected)", join.Comma(o.Scopes), count)
}
