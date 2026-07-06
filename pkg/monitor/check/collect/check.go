package collect

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Check(
	dryRun bool,
	parallel bool,
) {
	if parallel {
		collectors := fetch.List()
		fmt.Printf("Collectors: %s\n", join.Comma(collectors))

		return
	}

	for _, s := range fetch.List() {
		fmt.Printf("Command: %s\n", s)

		if dryRun {
			continue
		}

		if items := fetch.Run(s); len(items) > 0 {
			for _, i := range items {
				fmt.Printf("  %s: %s\n", i.Identifier, i.Detail)
			}
		} else {
			fmt.Println("  No items")
		}
	}
}
