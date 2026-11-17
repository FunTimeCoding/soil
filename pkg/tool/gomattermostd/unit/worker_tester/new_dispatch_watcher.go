package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"testing"
)

func NewDispatchWatcher(t *testing.T) *Watcher {
	t.Helper()
	s := NewStore(t)
	r := memory.New()
	chat := mock_client.New(t, "selfuser")
	result := &Watcher{
		Watcher: watcher.New(
			chat,
			NewGuard(t, chat, r),
			s,
			logger.New(t.Context()),
			r,
		),
		Store:  s,
		Client: chat,
	}
	t.Cleanup(result.Watcher.Close)

	return result
}
