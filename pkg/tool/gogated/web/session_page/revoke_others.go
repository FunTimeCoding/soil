package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"maragu.dev/gomponents"
)

func RevokeOthers() gomponents.Node {
	return action.New(
		constant.SessionsRevokeOthersPath,
		"Revoke all other sessions",
	).
		ButtonClass("secondary").
		Confirm("Sign out every other browser?").
		Node()
}
