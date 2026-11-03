package imports

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
)

func constantWant(
	e *import_entry.Entry,
	keeper string,
) string {
	if e.Path == keeper {
		return constant.ConstantPackage
	}

	return ConstantAlias(e.Path, e.Module)
}
