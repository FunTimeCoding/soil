package monitor

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (m *Monitor) refreshCache() {
	var all []*model.ChannelWithTeamData

	for offset := 0; ; offset += constant.MaxPerPage {
		page := m.client.MustAllChannels(constant.MaxPerPage, offset)
		all = append(all, page...)

		if len(page) < constant.MaxPerPage {
			break
		}
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, c := range all {
		m.channelCache[c.Name] = &c.Channel
	}
}
