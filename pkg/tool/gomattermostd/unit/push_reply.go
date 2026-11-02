package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/cross_service_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func pushReply(
	r *cross_service_tester.Tester,
	asked <-chan struct{},
) {
	r.Upstream.Push(
		postEvent(
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
