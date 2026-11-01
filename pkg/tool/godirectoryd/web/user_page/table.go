package user_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/user"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Table(users []*user.User) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(users))

	for _, u := range users {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(u.Account)),
				html.Td(gomponents.Text(u.Name)),
				html.Td(gomponents.Text(u.Mail)),
				html.Td(
					action.New(constant.DeleteUserPath, "Delete").
						Hidden(constant.AccountField, u.Account).
						Confirm(fmt.Sprintf("Delete user %s?", u.Account)).
						Small().
						Node(),
				),
			),
		)
	}

	return table.New("Account", "Name", "Mail", "").
		Rows(rows).
		Empty("No users yet.").
		Node()
}
