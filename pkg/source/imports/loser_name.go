package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
	"github.com/funtimecoding/soil/pkg/source/types/import_view"
)

func loserName(
	v *import_view.View,
	declaredName string,
	group []*import_entry.Entry,
	chosen map[string]bool,
	e *import_entry.Entry,
) string {
	taken := func(c string) bool {
		return c == declaredName || refusal(v, group, chosen, c) != ""
	}

	if declaredName != constant.MockPackage {
		return ChooseAlias(e.Path, e.Module, declaredName, taken)
	}

	if want := MockAlias(e.Path, e.Module); want != "" && !taken(want) {
		return want
	}

	return ""
}
