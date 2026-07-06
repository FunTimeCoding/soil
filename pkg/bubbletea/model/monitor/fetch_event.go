package monitor

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"time"
)

func (m *Model) fetchEvent(g fetch.Message) {
	m.items = g.Items
	rows := g.Rows()

	if !m.connect {
		for i := range rows {
			// Delete last column, user
			rows[i] = rows[i][:len(rows[i])-1]
		}
	}

	m.lastFetch = time.Now()
	m.table.SetRows(rows)
	m.updateColumns()
}
