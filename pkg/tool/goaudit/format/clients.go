package format

import "github.com/funtimecoding/soil/pkg/tool/goaudit/scan"

func Clients(clients []*scan.Client) string {
	t := newTable(
		[]string{
			"CLIENT",
			"REPO",
			"MUST",
			"BASIC",
			"ENTITY",
			"CONST",
			"EXAMPLE",
		},
	)

	for _, c := range clients {
		t.addRow(
			[]string{
				c.Path,
				c.Repo,
				mark(c.Must),
				mark(c.Basic),
				mark(c.Entity),
				mark(c.Constant),
				mark(c.Example),
			},
		)
	}

	return t.render()
}
