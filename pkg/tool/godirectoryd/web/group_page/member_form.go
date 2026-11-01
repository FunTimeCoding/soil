package group_page

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"maragu.dev/gomponents"
	"net/url"
)

func MemberForm(
	values url.Values,
	alert string,
) gomponents.Node {
	return entry.New(constant.AddMemberPath, "Add").
		Identifier(constant.MemberForm).
		Heading("Membership").
		Alert(alert).
		Text(constant.GroupField, "Group", values.Get(constant.GroupField)).
		Required().
		Text(
			constant.AccountField,
			"Account",
			values.Get(constant.AccountField),
		).
		Required().
		Alternative(constant.RemoveMemberPath, "Remove").
		Node()
}
