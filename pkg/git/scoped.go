package git

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func scoped(
	file string,
	scope string,
) (string, bool) {
	if scope == soil.CurrentDirectory {
		return file, true
	}

	return strings.CutPrefix(file, join.Empty(scope, stringsConstant.Slash))
}
