package gofix

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/element_format"
)

func fieldPadding(
	node dst.Node,
	parentLit *dst.CompositeLit,
) int {
	for _, el := range parentLit.Elts {
		kv, okay := el.(*dst.KeyValueExpr)

		if !okay {
			continue
		}

		if kv.Value != node && !containsNode(kv.Value, node) {
			continue
		}

		return element_format.AlignmentPadding(
			parentLit,
			element_format.KeyWidth(kv.Key),
		)
	}

	return 0
}
