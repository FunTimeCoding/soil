package example_table

import (
	"charm.land/bubbletea/v2"
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
)

func (m *Model) View() tea.View {
	return tea.NewView(
		fmt.Sprintf("%s%s", constant.Table.Render(m.table.View()), strings.Unix),
	)
}
