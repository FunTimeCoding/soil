package note

import "fmt"

func (n *Note) Format() string {
	return fmt.Sprintf("  %d: %s", n.Identifier, n.Text)
}
