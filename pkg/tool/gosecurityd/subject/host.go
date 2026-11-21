package subject

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
)

func Host(
	kind string,
	name string,
) string {
	return join.Empty(kind, constant.SubjectSeparator, name)
}
