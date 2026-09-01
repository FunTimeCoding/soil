package collect

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Check(
	dryRun bool,
	parallel bool,
) {
	if parallel {
		collectors := fetch.List()
		console.Format("Collectors: %s\n", join.Comma(collectors))

		return
	}

	for _, s := range fetch.List() {
		console.Format("Command: %s\n", s)

		if dryRun {
			continue
		}

		if items := fetch.Run(s); len(items) > 0 {
			for _, i := range items {
				console.Format("  %s: %s\n", i.Identifier, i.Detail)
			}
		} else {
			console.Line("  No items")
		}
	}
}
