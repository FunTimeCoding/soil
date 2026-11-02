package monitor

import (
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/mattermost/mattermost/server/public/model"
)

func (m *Monitor) refreshCache() error {
	var all []*model.ChannelWithTeamData

	for offset := 0; ; offset += chatConstant.MattermostMaxPerPage {
		page, e := m.client.AllChannels(
			chatConstant.MattermostMaxPerPage,
			offset,
		)

		if e != nil {
			return e
		}

		all = append(all, page...)

		if len(page) < chatConstant.MattermostMaxPerPage {
			break
		}
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	for _, c := range all {
		m.channelCache[c.Name] = &c.Channel
	}

	return nil
}
