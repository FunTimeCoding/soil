package monitor

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"time"
)

func (m *Monitor) processChannel(name string) error {
	m.mutex.RLock()
	c, okay := m.channelCache[name]
	lastMilli := m.lastCheckMillisecond[name]
	m.mutex.RUnlock()

	if !okay {
		return not_found.New("channel", name)
	}

	posts, e := m.client.PostsSince(c, time.UnixMilli(lastMilli))

	if e != nil {
		return e
	}

	m.mutex.Lock()
	m.lastCheckMillisecond[name] = nowMilli()
	m.mutex.Unlock()
	relevant := findRelevantPosts(posts, m.configuration.Topics)

	if len(relevant) == 0 {
		return nil
	}

	message := buildNotification(relevant, name, m.username)

	if message == "" {
		return nil
	}

	if _, f := m.client.PostSimple(m.notifyChannel, message); f != nil {
		return f
	}

	m.logger.Structured(
		"channel_notify",
		"channel",
		name,
		"count",
		len(relevant),
	)

	return nil
}
