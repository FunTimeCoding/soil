package reflow

import (
	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/text"
)

func linkSpan(
	source []byte,
	n ast.Node,
) (int, int, bool) {
	var destination text.SingleLineValue
	var title text.MultiLineValue

	switch l := n.(type) {
	case *ast.Link:
		if l.Reference != nil {
			return 0, 0, false
		}

		destination = l.Destination
		title = l.Title
	case *ast.Image:
		if l.Reference != nil {
			return 0, 0, false
		}

		destination = l.Destination
		title = l.Title
	default:
		return 0, 0, false
	}

	start := n.Pos()

	if start < 0 || destination.IsOwned() || destination.IsEmpty() {
		return 0, 0, false
	}

	stop := destination.Index().Stop

	if indices := title.Indices(); !title.IsOwned() && len(indices) > 0 {
		stop = indices[len(indices)-1].Stop
	}

	for stop < len(source) && source[stop] != ')' {
		stop++
	}

	if stop >= len(source) {
		return 0, 0, false
	}

	return start, stop + 1, true
}
