package mattermost_notifier

import "github.com/funtimecoding/soil/pkg/chat/mattermost"

func New(
	m *mattermost.Client,
	channel string,
	prefix string,
) *Notifier {
	h := m.MustTeamChannel(channel)

	return &Notifier{mattermost: m, channel: h, prefix: prefix}
}
