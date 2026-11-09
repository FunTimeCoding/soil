package monitor

import "github.com/funtimecoding/soil/pkg/monitor/item"

func (m *Model) selectedItem() *item.Item {
	r := m.table.SelectedRow()

	if r == nil {
		return nil
	}

	return m.labelItem(r[0])
}
