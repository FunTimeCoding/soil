package browser_tester

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/accessibility"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (b *Browser) Snapshot() []*SnapshotNode {
	b.T.Helper()
	var nodes []*accessibility.Node
	errors.PanicOnError(
		chromedp.Run(
			b.Context,
			chromedp.ActionFunc(
				func(c context.Context) error {
					var e error
					nodes, e = accessibility.GetFullAXTree().Do(c)

					return e
				},
			),
		),
	)
	lookup := make(map[accessibility.NodeID]*SnapshotNode)
	var roots []*SnapshotNode
	uid := 0

	for _, n := range nodes {
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
		sn := &SnapshotNode{
			UID:   fmt.Sprintf("e%d", uid),
			Role:  role,
			Name:  name,
			Value: value,
		}
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
