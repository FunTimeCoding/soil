package symbol

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func memberName(
	receiver string,
	name string,
) string {
	if receiver == "" {
		return name
	}

	return join.Empty(receiver, stringsConstant.Dot, name)
}
