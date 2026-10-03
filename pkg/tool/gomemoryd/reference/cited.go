package reference

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strings"
)

func cited(
	span string,
	named func(
		scope string,
		name string,
	) bool,
) bool {
	scope, name, found := strings.Cut(
		strings.TrimPrefix(span, constant.MemoryScheme),
		stringsConstant.Slash,
	)

	return found && name != "" && named(scope, name)
}
