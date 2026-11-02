package monitor

func (m *Monitor) initialize() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.initialized {
		return nil
	}

	username := m.configuration.Username

	if username == "" {
		me, e := m.client.Me()

		if e != nil {
			return e
		}

		username = me.Username
	}

	name := m.configuration.NotificationChannel

	if name == "" {
		name = "town-square"
	}

	channel, e := m.client.TeamChannel(name)

	if e != nil {
		return e
	}

	m.username = username
	m.notifyChannel = channel
	now := nowMilli()

	for _, c := range m.configuration.Channels {
		m.lastCheckMillisecond[c] = now
	}

	m.initialized = true
	m.logger.Structured(
		"monitor_initialize",
		"username",
		m.username,
		"notify_channel",
		m.notifyChannel.Name,
	)

	return nil
}
