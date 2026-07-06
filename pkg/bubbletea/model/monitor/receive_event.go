package monitor

import (
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/receive"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"log"
)

func (m *Model) receiveEvent(s receive.Message) {
	for _, line := range s.Text {
		p := split.Comma(line)
		log.Printf("receive event: %s\n", line)

		switch p[0] {
		case constant.FlagAddCommand:
			handle := p[1]
			identifier := p[2]
			rows := m.table.Rows()
			rows[rowIndex(m.table, identifier)][2] = handle
			m.table.SetRows(rows)
			m.updateColumns()
		case constant.FlagRemoveCommand:
			handle := p[1]
			identifier := p[2]
			rows := m.table.Rows()
			i := rowIndex(m.table, identifier)

			if i != -1 {
				row := rows[i]
				log.Printf(
					"remove handle %s from %s\n",
					handle,
					row[0],
				)

				if row[2] == handle {
					row[2] = ""
					m.table.SetRows(rows)
					m.updateColumns()
				}
			}
		}
	}
}
