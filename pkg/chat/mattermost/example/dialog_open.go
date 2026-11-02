package example

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/mattermost/mattermost/server/public/model"
)

func dialogOpen(m *mattermost.Client) {
	m.MustOpenDialog(
		model.OpenDialogRequest{
			TriggerId: "",
			URL: locator.New(constant.Localhost).Port(
				constant.ListenPort,
			).Insecure().Path(
				constant.DialogPath,
			).String(),
			Dialog: model.Dialog{
				CallbackId:       "myOwnId",
				Title:            "Maximum 24 characters",
				IntroductionText: "Markdown formatted text",
				IconURL:          "",
				Elements:         nil,
				SubmitLabel:      "",
				NotifyOnCancel:   false,
				State:            "",
			},
		},
	)
}
