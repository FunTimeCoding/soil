package monitor

func (m *Model) Close() {
	if m.hold != nil {
		m.hold.Close()
	}

	if m.monitor != nil {
		m.monitor.Close()
	}
}
