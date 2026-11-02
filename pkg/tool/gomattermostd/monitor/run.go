package monitor

import "context"

func (m *Monitor) run() error {
	return m.guard.Attempt(
		context.Background(),
		func(context.Context) error {
			if e := m.initialize(); e != nil {
				return e
			}

			if e := m.refreshCache(); e != nil {
				return e
			}

			for _, name := range m.configuration.Channels {
				if f := m.processChannel(name); f != nil {
					return f
				}
			}

			return nil
		},
	)
}
