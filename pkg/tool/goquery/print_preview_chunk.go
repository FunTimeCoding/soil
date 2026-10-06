package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"strings"
)

func printPreviewChunk(
	k client.PreviewChunk,
	allowance int,
) int {
	console.Format(
		"  chunk %d  lines %d-%d  %d B  %d tokens",
		k.Index,
		k.FirstLine,
		k.LastLine,
		k.Bytes,
		k.Tokens,
	)

	if k.Piece {
		console.Format("  table piece")
	}

	if k.Tokens <= allowance {
		console.Line()

		return 0
	}

	console.Format("  OVER by %d\n", k.Tokens-allowance)

	switch {
	case k.CutLevel > 0:
		console.Format(
			"      fixable: a %s heading (or shallower) before line %s makes this window fit\n",
			strings.Repeat("#", k.CutLevel),
			lineRanges(k.CutLines),
		)
	case k.CutChecked:
		console.Format("      no single heading makes this window fit\n")
	}

	return 1
}
