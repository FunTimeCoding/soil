package seed_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Priority(seeds []seed.Seed) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(seeds))

	for _, v := range seeds {
		rows = append(
			rows,
			html.Tr(
				html.Td(
					html.Class(constant.DragHandleClass),
					gomponents.Text("⠿"),
					html.Input(
						html.Type("hidden"),
						html.Name(constant.ItemParameter),
						html.Value(fmt.Sprintf("%d", v.Identifier)),
					),
				),
				html.Td(
					html.Class(webConstant.NumericClass),
					gomponents.Textf("%d", v.Position),
				),
				html.Td(gomponents.Text(v.Name)),
				html.Td(
					html.Class(constant.SeedPathClass),
					gomponents.Text(seedDirectory(v.Path)),
				),
			),
		)
	}

	return table.New("", "#", "Name", "Path").
		Numeric(1).
		Class(constant.SeedTableClass).
		BodyIdentifier(constant.SeedOrderMark).
		Rows(rows).
		Empty("No seeds.").
		Node()
}
