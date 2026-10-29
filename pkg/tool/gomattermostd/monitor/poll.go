package monitor

import "github.com/funtimecoding/soil/pkg/errors/classify"

func (m *Monitor) Poll() {
	e := m.run()

	if e == nil {
		return
	}

	if classify.Reportable(e) {
		m.reporter.CaptureException(e)

		return
	}

	m.logger.Plain("monitor poll: %v", e)
}
