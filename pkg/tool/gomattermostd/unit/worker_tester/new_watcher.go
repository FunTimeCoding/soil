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

func NewWatcher(
	t *testing.T,
	window time.Duration,
) *Watcher {
	t.Helper()
	sink, c := NewSink(t)
	s := NewStore(t)
	r := memory.New()

	return &Watcher{
		Watcher: watcher.New(
			mock_client.New("selfuser"),
			s,
			notifier.New(c, "mattermost", r),
			logger.New(context.Background()),
			r,
			window,
		),
		Store: s,
		Sink:  sink,
	}
}
