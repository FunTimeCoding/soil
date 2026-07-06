package format

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
	"strings"
)

func Services(services []*scan.Service) string {
	var b strings.Builder
	t := newTable(
		[]string{
			"SERVICE",
			"REPO",
			"MCP",
			"REST",
			"WEB",
			"STORE",
			"GEN",
			"CVT",
			"CLI",
			"TYPES",
			"MODEL",
			"CONST",
			"WORK",
			"TEST",
			"OPT",
			"RUN",
		},
	)

	for _, s := range services {
		t.addRow(
			[]string{
				s.Name,
				s.Repo,
				mark(s.ModelContext),
				mark(s.Server),
				mark(s.Web),
				mark(s.Store),
				mark(s.Generated),
				mark(s.Convert),
				mark(s.Client),
				mark(s.Types),
				mark(s.Model),
				constantMark(s),
				mark(s.Worker),
				mark(s.IntegrationTests),
				mark(s.Option),
				mark(s.Run),
			},
		)
	}

	b.WriteString(t.render())
	warnings := collectWarnings(services)

	if len(warnings) > 0 {
		b.WriteByte('\n')

		for _, w := range warnings {
			b.WriteString(w)
			b.WriteByte('\n')
		}
	}

	return b.String()
}
