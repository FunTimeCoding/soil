package monitor

func (m *Monitor) Close() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if !m.running {
		return
	}

	m.scheduler.Close()
	m.running = false
}
