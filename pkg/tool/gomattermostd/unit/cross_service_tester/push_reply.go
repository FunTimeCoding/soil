package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/event_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (o *Tester) PushReply(asked <-chan struct{}) {
	o.Upstream.Push(
		event_tester.Posted(
			&model.Post{
				Id:       "reply",
				RootId:   "alfa",
				UserId:   "foxtrot",
				Message:  "first message",
				CreateAt: time.Now().UnixMilli(),
			},
		),
	)
	<-asked
}
