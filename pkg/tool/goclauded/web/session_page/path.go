package session_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/url"
)

func Path(identifier string) string {
	return join.Empty(
		constant.SessionsPath,
		stringsConstant.Slash,
		url.PathEscape(identifier),
	)
}
