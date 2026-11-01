package group_page

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"net/url"
)

func CreateForm(
	values url.Values,
	alert string,
) gomponents.Node {
	return entry.New(constant.CreateGroupPath, "Create").
		Identifier(constant.CreateGroupForm).
		Heading("Create group").
		Alert(alert).
		Text(constant.NameField, "Name", values.Get(constant.NameField)).
		Required().
		Node()
}
