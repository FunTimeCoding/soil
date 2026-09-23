package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
	"testing"
	"time"
)

func NewDispatchWatcher(t *testing.T) *Watcher {
	t.Helper()
	sink, c := NewSink(t)
	s := NewStore(t)
	r := memory.New()
	chat := mock_client.New("selfuser")

	return &Watcher{
		Watcher: watcher.New(
			chat,
			s,
			notifier.New(c, "mattermost", r),
			logger.New(context.Background()),
			r,
			time.Hour,
		),
		Store:  s,
		Client: chat,
		Sink:   sink,
	}
}
