package loki

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func printOverview(
	entries []*overview,
	f *option.Format,
) {
	for _, e := range entries {
		s := status.New(f).String(
			formatNamespace(e, f),
		).Integer(e.Count).String(
			formatLatest(e),
		)
		console.Line(s.Format())
	}
}
