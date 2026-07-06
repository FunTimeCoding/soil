package monitor

import (
	"charm.land/bubbletea/v2"
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/receive"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/tick"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/time"
)

func (m *Model) tickEvent(g tick.Message) (*Model, tea.Cmd) {
	var result tea.BatchMsg

	if m.connect && m.second == 0 {
		m.client.Connect()
		m.client.Write(
			join.Comma([]string{constant.LoginCommand, m.user, m.hostname}),
		)
	}

	if m.second%60 == 0 {
		if m.second == 0 {
			// Load on startup
			result = append(result, fetch.Command())
		} else {
			if m.connect {
				m.client.Write(constant.PingCommand)
			}

			if m.auto {
				result = append(result, fetch.Command())
			}
		}
	}

	f := option.ExtendedColor.Copy()
	top := status.New(f).String()
	top.String(fmt.Sprintf("items: %d", len(m.table.Rows())))

	if m.auto {
		top.String("auto")
	} else {
		top.String("manual")
		top.String(
			fmt.Sprintf(
				"last fetch: %s",
				m.lastFetch.Format(time.DateMinute),
			),
		)
	}

	m.topBar = top.Format()
	bottom := status.New(f)
	bottom.String(m.hostname)

	if false {
		bottom.String(g.Time.Format(time.DateSecond))
	}

	if false {
		bottom.Integer(60 - m.second%60)
	}

	bottom.String(fmt.Sprintf("%dx%d", m.width, m.height))
	m.bottomBar = bottom.Format()
	m.second++
	result = append(result, tick.Command())

	if m.connect {
		result = append(result, receive.Command(m.client))
	}

	return m, tea.Batch(result...)
}
