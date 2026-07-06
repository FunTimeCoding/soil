package monitor

import "github.com/funtimecoding/soil/pkg/errors/panics"

func (m *Monitor) processChannel(name string) {
	m.mutex.RLock()
	c, okay := m.channelCache[name]
	lastMilli := m.lastCheckMillisecond[name]
	m.mutex.RUnlock()

	if !okay {
		panics.Print("name not found in channel cache: %s", name)
	}

	posts := m.client.MustRecentPosts(c, lastMilli)
	m.mutex.Lock()
	m.lastCheckMillisecond[name] = nowMilli()
	m.mutex.Unlock()
	relevant := findRelevantPosts(posts, m.configuration.Topics)

	if len(relevant) == 0 {
		return
	}

	message := buildNotification(relevant, name, m.username)

	if message == "" {
		return
	}

	m.client.MustPostSimple(m.notifyChannel, message)
	m.logger.Structured(
		"channel_notify",
		"channel",
		name,
		"count",
		len(relevant),
	)
}
