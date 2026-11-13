package conversation_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/url"
	"strconv"
)

func sidebarPath(offset int) string {
	path := join.Empty(
		constant.ConversationsPath,
		stringsConstant.Slash,
		constant.SidebarAction,
	)

	if offset == 0 {
		return path
	}

	return join.Empty(
		path,
		stringsConstant.Question,
		url.Values{constant.Offset: {strconv.Itoa(offset)}}.Encode(),
	)
}
