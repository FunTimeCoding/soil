package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
)

func (o *Tester) DispatchWatcher() *Watcher {
	o.t.Helper()
	s := o.Store()
	r := memory.New()
	chat := mock_client.New(o.t, "selfuser")
	result := &Watcher{
		Watcher: watcher.New(
			chat,
			o.Guard(chat, r),
			s,
			logger.New(o.t.Context()),
			r,
		),
		Store:  s,
		Client: chat,
	}
	o.t.Cleanup(result.Watcher.Close)

	return result
}
