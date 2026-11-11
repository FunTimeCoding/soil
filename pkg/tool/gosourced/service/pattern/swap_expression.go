package pattern

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/dstutil"
)

func swapExpression(
	file dst.Node,
	old dst.Node,
	replacement dst.Expr,
) error {
	found := false
	dstutil.Apply(
		file,
		func(c *dstutil.Cursor) bool {
			if c.Node() != old {
				return true
			}

			c.Replace(replacement)
			found = true

			return false
		},
		nil,
	)

	if !found {
		return fmt.Errorf("expression not found in its file")
	}

	return nil
}
