package service

import (
	"github.com/dave/dst"
	"github.com/dave/dst/dstutil"
)

func hasDecorations(node dst.Node) bool {
	found := false
	dstutil.Apply(
		node,
		func(c *dstutil.Cursor) bool {
			n := c.Node()

			if n == nil {
				return false
			}

			decorations := n.Decorations()

			if len(decorations.Start.All()) > 0 ||
				len(decorations.End.All()) > 0 {
				found = true
			}

			return !found
		},
		nil,
	)

	return found
}
