package subscribed

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"github.com/mattermost/mattermost/server/public/model"
	"time"
)

func (o *Subscribed) Reply(
	identifier string,
	root string,
	message string,
) {
	o.Watcher.Dispatch(
		worker_tester.PostEvent(
			&model.Post{
				Id:       identifier,
				RootId:   root,
				UserId:   "foxtrot",
				Message:  message,
				CreateAt: time.Now().UnixMilli(),
			},
		),
	)
}
