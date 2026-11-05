package goquery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"strings"
)

func printPreviewSections(
	t *terminal.Terminal,
	sections []client.PreviewSection,
	allowance int,
) {
	if len(sections) == 0 {
		return
	}

	t.Line("sections")

	for _, x := range sections {
		line := fmt.Sprintf(
			"  %s%-*s lines %d-%d  %d tokens",
			strings.Repeat("  ", max(x.Level-1, 0)),
			48-2*max(x.Level-1, 0),
			label(x.Title),
			x.FirstLine,
			x.LastLine,
			x.Tokens,
		)

		if x.Tokens <= allowance {
			t.Line(line)

			continue
		}

		t.Line(
			join.Empty(line, fmt.Sprintf("  OVER by %d", x.Tokens-allowance)),
		)

		for _, b := range x.Blocks {
			t.Line(
				fmt.Sprintf(
					"      %-12s lines %d-%d  %d tokens",
					b.Kind,
					b.FirstLine,
					b.LastLine,
					b.Tokens,
				),
			)
		}
	}

	t.Line()
}
