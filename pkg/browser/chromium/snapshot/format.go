package snapshot

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
		line := fmt.Sprintf(
			"%suid=%s [%s] %s",
			strings.Repeat("  ", depth),
			n.UID,
			n.Role,
			n.Name,
		)

		if n.Value != "" {
			line = fmt.Sprintf("%s = %s", line, n.Value)
		}

		b.WriteString(line)
		b.WriteByte('\n')

		if len(n.Children) > 0 {
			b.WriteString(Format(n.Children, depth+1))
		}
	}

	return b.String()
}
