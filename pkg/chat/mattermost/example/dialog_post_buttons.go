package example

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
)

func dialogPostButtons(m *mattermost.Client) {
	p := &model.Post{
		ChannelId: m.DefaultChannel().Id,
		Message:   "Some message with buttons",
	}
	callback := locator.New(constant.Localhost).Port(
		constant.ListenPort,
	).Insecure().Path(
		constant.CallbackPath,
	).String()
	model.ParseMessageAttachment(
		p,
		[]*model.MessageAttachment{
			{
				AuthorName: "alfa",
				Title:      "bravo",
				Text:       "charlie",
				Color:      "#36a64f",
				Fields: []*model.MessageAttachmentField{
					{Title: "Current Version", Value: "v2.1.0", Short: true},
					{Title: "Target Version", Value: "v2.2.1", Short: true},
				},
				Actions: []*model.PostAction{
					{
						Id:    "yes",
						Name:  "Yes",
						Type:  model.PostActionTypeButton,
						Style: "primary",
						Integration: &model.PostActionIntegration{
							Context: model.StringInterface{
								"action":  "update_server",
								"server":  "prod-web-01",
								"version": "v2.2.1",
							},
							URL: callback,
						},
					},
					{
						Id:    "cancel",
						Name:  "Cancel",
						Type:  model.PostActionTypeButton,
						Style: "danger",
						Integration: &model.PostActionIntegration{
							Context: model.StringInterface{
								"action":  "cancel_update",
								"server":  "prod-web-01",
								"version": "v2.2.1",
							},
							URL: callback,
						},
					},
				},
			},
		},
	)
	m.MustPost(p)
}
