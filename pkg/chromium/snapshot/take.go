package snapshot

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/accessibility"
)

func Take(c context.Context) ([]*Node, error) {
	nodes, e := accessibility.GetFullAXTree().Do(c)

	if e != nil {
		return nil, e
	}

	lookup := make(map[accessibility.NodeID]*Node)
	var roots []*Node
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
		sn := &Node{
			UID:                      fmt.Sprintf("e%d", uid),
			Role:                     role,
			Name:                     name,
			Value:                    value,
			BackendDOMNodeIdentifier: int64(n.BackendDOMNodeID),
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

	return roots, nil
}
