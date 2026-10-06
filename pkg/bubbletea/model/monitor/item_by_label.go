package monitor

import "github.com/funtimecoding/soil/pkg/monitor/item"

func (m *Model) itemByLabel(label string) *item.Item {
	for _, i := range m.items {
		if i.Label == label {
			return i
		}
	}

	return nil
}
