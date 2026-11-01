package seed_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/seed"
	"github.com/funtimecoding/soil/pkg/web/component/table"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Recent(seeds []seed.Seed) gomponents.Node {
	rows := make([]gomponents.Node, 0, len(seeds))

	for _, v := range seeds {
		rows = append(
			rows,
			html.Tr(
				layout.TimeCell(v.ModifiedAt),
				html.Td(gomponents.Text(v.Name)),
				html.Td(
					html.Class(constant.SeedPathClass),
					gomponents.Text(seedDirectory(v.Path)),
				),
			),
		)
	}

	return table.New("Modified", "Name", "Path").
		Class(constant.SeedTableClass).
		Rows(rows).
		Empty("No seeds.").
		Node()
}
