package example

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/mattermost/mattermost/server/public/model"
)

func dialogSubmit(m *mattermost.Client) {
	m.MustSubmitDialog(
		model.SubmitDialogRequest{
			Type:       "",
			URL:        "",
			CallbackId: "",
			State:      "",
			UserId:     "",
			ChannelId:  m.DefaultChannel().Id,
			TeamId:     m.DefaultTeam().Id,
			Submission: nil,
			Cancelled:  false,
		},
	)
}
