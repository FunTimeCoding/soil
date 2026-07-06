package monitor

import (
	"charm.land/lipgloss/v2"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/toast"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func renderToasts(v []*toast.Toast) string {
	if len(v) == 0 {
		return ""
	}

	var lines []string

	for _, n := range v {
		lines = append(lines, n.Format())
	}

	if len(lines) == 0 {
		return ""
	}

	return lipgloss.NewStyle().Foreground(
		lipgloss.Color("205"), // Magenta-ish
	).Render(join.NewLine(lines))
}
