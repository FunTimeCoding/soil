package goquery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"strings"
)

func printPreviewChunk(
	t *terminal.Terminal,
	k client.PreviewChunk,
	allowance int,
	embedAllowance int,
) int {
	line := fmt.Sprintf(
		"  chunk %d  lines %d-%d  %d B  %d tokens  embed %d",
		k.Index,
		k.FirstLine,
		k.LastLine,
		k.Bytes,
		k.Tokens,
		k.EmbedTokens,
	)

	if k.Piece {
		line = join.Empty(line, "  table piece")
	}

	over := k.Tokens > allowance
	embedOver := k.EmbedTokens > embedAllowance

	if over {
		line = join.Empty(line, fmt.Sprintf("  OVER by %d", k.Tokens-allowance))
	}

	t.Line(line)

	if !over && !embedOver {
		return 0
	}

	if embedOver {
		t.Line(
			fmt.Sprintf(
				"      embedder: over its %d tokens by %d, refused when embedded",
				embedAllowance,
				k.EmbedTokens-embedAllowance,
			),
		)
	}

	if !over {
		return 1
	}

	switch {
	case k.CutLevel > 0:
		t.Line(
			fmt.Sprintf(
				"      fixable: a %s heading (or shallower) before line %s makes this window fit",
				strings.Repeat("#", k.CutLevel),
				lineRanges(k.CutLines),
			),
		)
	case k.CutChecked:
		t.Line("      no single heading makes this window fit")
	}

	return 1
}
