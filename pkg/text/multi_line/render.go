package multi_line

import "github.com/funtimecoding/soil/pkg/strings/join"

func (m *MultiLine) Render() string {
	return join.NewLine(m.lines)
}
