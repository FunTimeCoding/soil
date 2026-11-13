package conversation_page

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func searchPath() string {
	return join.Empty(
		constant.ConversationsPath,
		stringsConstant.Slash,
		constant.SearchAction,
	)
}
