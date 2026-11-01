package group_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func Table(groups []*group.Group) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(groups))

	for _, g := range groups {
		rows = append(
			rows,
			html.Tr(
				html.Td(gomponents.Text(g.Name)),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Text(strconv.Itoa(g.Number)),
				),
				html.Td(gomponents.Text(join.CommaSpace(g.Member))),
				html.Td(
					action.New(constant.DeleteGroupPath, "Delete").
						Hidden(constant.NameField, g.Name).
						Confirm(fmt.Sprintf("Delete group %s?", g.Name)).
						Small().
						Node(),
				),
			),
		)
	}

	return table.New("Name", "Number", "Members", "").
		Numeric(1).
		Rows(rows).
		Empty("No groups yet.").
		Node()
}
