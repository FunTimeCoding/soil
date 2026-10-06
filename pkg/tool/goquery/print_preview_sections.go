package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"strings"
)

func printPreviewSections(
	sections []client.PreviewSection,
	allowance int,
) {
	if len(sections) == 0 {
		return
	}

	console.Format("sections\n")

	for _, x := range sections {
		console.Format(
			"  %s%-*s lines %d-%d  %d tokens",
			strings.Repeat("  ", max(x.Level-1, 0)),
			48-2*max(x.Level-1, 0),
			label(x.Title),
			x.FirstLine,
			x.LastLine,
			x.Tokens,
		)

		if x.Tokens <= allowance {
			console.Line()

			continue
		}

		console.Format("  OVER by %d\n", x.Tokens-allowance)

		for _, b := range x.Blocks {
			console.Format(
				"      %-12s lines %d-%d  %d tokens\n",
				b.Kind,
				b.FirstLine,
				b.LastLine,
				b.Tokens,
			)
		}
	}

	console.Line()
}
