package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"strings"
)

func printPreviewChunk(
	k client.PreviewChunk,
	allowance int,
	embedAllowance int,
) int {
	console.Format(
		"  chunk %d  lines %d-%d  %d B  %d tokens  embed %d",
		k.Index,
		k.FirstLine,
		k.LastLine,
		k.Bytes,
		k.Tokens,
		k.EmbedTokens,
	)

	if k.Piece {
		console.Format("  table piece")
	}

	over := k.Tokens > allowance
	embedOver := k.EmbedTokens > embedAllowance

	if !over && !embedOver {
		console.Line()

		return 0
	}

	if over {
		console.Format("  OVER by %d", k.Tokens-allowance)
	}

	console.Line()

	if embedOver {
		console.Format(
			"      embedder: over its %d tokens by %d, refused when embedded\n",
			embedAllowance,
			k.EmbedTokens-embedAllowance,
		)
	}

	if !over {
		return 1
	}

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
