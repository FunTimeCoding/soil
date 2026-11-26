package node

import (
	"fmt"
	"strings"
)

func Format(
	v []*Node,
	depth int,
) string {
	var b strings.Builder

	for _, n := range v {
		b.WriteString(strings.Repeat("  ", depth))

		if n.UID != "" {
			_, _ = fmt.Fprintf(&b, "uid=%s ", n.UID)
		}

		_, _ = fmt.Fprintf(&b, "[%s] %s", n.Role, n.Name)

		if n.Value != "" {
			_, _ = fmt.Fprintf(&b, " = %s", n.Value)
		}

		b.WriteByte('\n')
		b.WriteString(Format(n.Children, depth+1))
	}

	return b.String()
}
