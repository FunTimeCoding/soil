package session_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func ActionPath(
	identifier string,
	action string,
) string {
	return join.Empty(Path(identifier), stringsConstant.Slash, action)
}
