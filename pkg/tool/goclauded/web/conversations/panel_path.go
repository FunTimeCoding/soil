package conversations

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/url"
)

func panelPath(identifier string) string {
	return join.Empty(
		constant.ConversationsPath,
		stringsConstant.Slash,
		url.PathEscape(identifier),
	)
}
