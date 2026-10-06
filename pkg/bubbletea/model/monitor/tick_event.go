package monitor

import (
	"charm.land/bubbletea/v2"
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
	console "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
)

func (m *Model) tickEvent(_ tick.Message) (*Model, tea.Cmd) {
	var result tea.BatchMsg

	if m.second%60 == 0 && (m.second == 0 || m.auto) {
		result = append(result, fetch.Command())
	}

	f := console.ExtendedColorFormat.Copy()
	top := status.New(f).String()
	top.String(fmt.Sprintf("items: %d", len(m.table.Rows())))

	if m.auto {
		top.String("auto")
	} else {
		top.String("manual")
		top.String(
			fmt.Sprintf(
				"last fetch: %s",
				m.lastFetch.Format(timeConstant.DateMinute),
			),
		)
	}

	if m.monitor != nil {
		if m.claimError != nil {
			top.String("claims offline")
		} else {
			top.String(fmt.Sprintf("claims: %d", len(m.claims)))
		}
	}

	m.topBar = top.Format()
	bottom := status.New(f)
	bottom.String(m.hostname)
	bottom.String(fmt.Sprintf("%dx%d", m.width, m.height))
	m.bottomBar = bottom.Format()
	m.second++
	result = append(result, tick.Command())

	return m, tea.Batch(result...)
}
