package monitor

import "github.com/funtimecoding/soil/pkg/strings/join"

func (m *Model) applyClaims() {
	if !m.connect {
		return
	}

	rows := m.table.Rows()

	for _, r := range rows {
		if i := m.itemByLabel(r[0]); i != nil {
			r[len(r)-1] = join.Comma(m.claims[i.Identifier])
		}
	}

	m.table.SetRows(rows)
	m.updateColumns()
}
