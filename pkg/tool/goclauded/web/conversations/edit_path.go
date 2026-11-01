package conversations

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func editPath(identifier string) string {
	return join.Empty(
		panelPath(identifier),
		stringsConstant.Slash,
		constant.EditAction,
	)
}
