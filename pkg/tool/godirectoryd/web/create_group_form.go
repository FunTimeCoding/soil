package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func createGroupForm() gomponents.Node {
	return html.Form(
		html.Method("post"),
		html.Action(constant.CreateGroupPath),
		html.H2(gomponents.Text("Create group")),
		field(constant.NameField, "Name", constant.TextInputType, true),
		html.Button(html.Type("submit"), gomponents.Text("Create")),
	)
}
