package browser_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/types/snapshot_node"
	"strings"
)

func formatNodes(
	nodes []*snapshot_node.Node,
	depth int,
) string {
	var b strings.Builder
	indent := strings.Repeat("  ", depth)

	for _, n := range nodes {
		line := fmt.Sprintf("%s[%s] %s", indent, n.Role, n.Name)

		if n.Value != "" {
			line = fmt.Sprintf("%s = %s", line, n.Value)
		}

		b.WriteString(line)
		b.WriteByte('\n')

		if len(n.Children) > 0 {
			b.WriteString(formatNodes(n.Children, depth+1))
		}
	}

	return b.String()
}
