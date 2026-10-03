package monitor

import "github.com/funtimecoding/soil/pkg/bubbletea/constant"

func (m *Model) updateTableHeight(
	addOne bool,
	removeOne bool,
) {
	static := constant.MonitorTopBarHeight +
		constant.MonitorBottomBarHeight +
		constant.MonitorTableHeaderHeight
	toasts := len(m.toast)

	if toasts > 1 {
		static += toasts - 1
	}

	height := m.height - static

	if addOne {
		height += 1
	}

	if removeOne {
		height -= 1
	}

	m.table.SetHeight(height)
}
