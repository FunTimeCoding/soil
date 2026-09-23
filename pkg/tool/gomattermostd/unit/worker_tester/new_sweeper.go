package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_indexer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/notifier"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/sweeper"
	"testing"
)

func NewSweeper(t *testing.T) *Sweeper {
	t.Helper()
	sink, c := NewSink(t)
	s := NewStore(t)
	r := memory.New()
	index := mock_indexer.New()

	return &Sweeper{
		Sweeper: sweeper.New(
			s,
			notifier.New(c, "mattermost", r),
			index,
			logger.New(context.Background()),
			r,
			constant.PurgeWindow,
			constant.PurgeInterval,
		),
		Store: s,
		Index: index,
		Sink:  sink,
	}
}
