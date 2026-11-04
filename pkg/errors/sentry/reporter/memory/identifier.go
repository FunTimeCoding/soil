package memory

import "fmt"

func (m *Memory) identifier() string {
	return fmt.Sprintf("memory-%d", len(m.events))
}
