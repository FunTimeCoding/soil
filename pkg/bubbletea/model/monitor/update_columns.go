package monitor

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/constant"
	"log"
)

func (m *Model) updateColumns() {
	if m.width == 0 {
		return
	}

	remaining := m.width - 2 // 2 for border
	var detailIndex int

	for i, c := range m.table.Columns() {
		switch c.Title {
		case constant.ItemIdentifierColumn:
			w := columnWidth(c, m.table, i)
			m.table.Columns()[i].Width = w
			remaining -= w + 2 // 2 for padding
		case constant.ItemScoreColumn:
			w := columnWidth(c, m.table, i)
			m.table.Columns()[i].Width = w
			remaining -= w + 2 // 2 for padding
		case constant.ItemSeverityColumn:
			w := columnWidth(c, m.table, i)
			m.table.Columns()[i].Width = w
			remaining -= w + 2 // 2 for padding
		case constant.ItemDetailColumn:
			detailIndex = i
		case constant.ItemUserColumn:
			w := columnWidth(c, m.table, i)
			m.table.Columns()[i].Width = w
			remaining -= w + 2 // 2 for padding
		default:
			log.Panicf("unexpected: %s", c.Title)
		}
	}

	m.table.Columns()[detailIndex].Width = remaining - 2 // 2 for padding
	m.table.SetWidth(m.width)
}
