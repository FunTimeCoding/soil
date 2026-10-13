package monitor

func (m *Model) Close() {
	if m.monitor != nil {
		m.monitor.Close()
	}
}
