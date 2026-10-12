package imports

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
)

func refusal(
	v *import_view.View,
	group []*import_entry.Entry,
	chosen map[string]bool,
	name string,
) string {
	if chosen[name] {
		return fmt.Sprintf("%s already given in this file", name)
	}

	if isBuiltin(name) {
		return "builtin collision"
	}

	if v.LocalTaken != nil && v.LocalTaken(name) {
		return fmt.Sprintf("local collision with %q", name)
	}

	if h := holder(v, group, name); h != "" {
		return fmt.Sprintf("import collision with %s", h)
	}

	return ""
}
