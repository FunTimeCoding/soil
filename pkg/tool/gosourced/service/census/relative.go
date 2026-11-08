package census

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"strings"
)

func relative(
	module string,
	path string,
) string {
	path = strings.TrimSuffix(path, constant.TestClauseSuffix)

	if path == module {
		return constant.DotName
	}

	return strings.TrimPrefix(path, join.Empty(module, "/"))
}
