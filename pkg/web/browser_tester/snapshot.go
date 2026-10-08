package browser_tester

import (
	"fmt"
	"github.com/chromedp/cdproto/accessibility"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/types/snapshot_node"
)

func (b *Browser) Snapshot() []*snapshot_node.Node {
	b.T.Helper()
	tree, e := chromedp.Call(
		b.Context,
		accessibility.GetFullAXTree,
		accessibility.GetFullAXTreeParams{},
	)
	errors.PanicOnError(e)
	lookup := make(map[accessibility.NodeID]*snapshot_node.Node)
	var roots []*snapshot_node.Node
	uid := 0

	for _, n := range tree.Nodes {
		if n.Ignored {
			continue
		}

		role := ""

		if n.Role != nil {
			role = fmt.Sprintf("%v", n.Role.Value)
		}

		name := ""

		if n.Name != nil {
			name = fmt.Sprintf("%v", n.Name.Value)
		}

		value := ""

		if n.Value != nil {
			value = fmt.Sprintf("%v", n.Value.Value)
		}

		uid++
		sn := snapshot_node.New(fmt.Sprintf("e%d", uid), role, name, value)
		lookup[n.NodeID] = sn

		if n.ParentID != "" {
			if parent, okay := lookup[n.ParentID]; okay {
				parent.Children = append(parent.Children, sn)

				continue
			}
		}

		roots = append(roots, sn)
	}

	return roots
}
