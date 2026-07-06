package monitor

import "github.com/funtimecoding/soil/pkg/monitor/item"

func (m *Model) selectedItem() *item.Item {
	identifier := m.table.SelectedRow()[0]

	for _, i := range m.items {
		if i.Identifier == identifier {
			return i
		}
	}

	return nil
}
