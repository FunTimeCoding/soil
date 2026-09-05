package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"testing"
	"time"
)

func NewWithStore(
	t *testing.T,
	s *store.Store,
	c *mock_client.Client,
) *Tester {
	t.Helper()

	return &Tester{
		Store: s,
		Worker: worker.New(
			c,
			s,
			logger.New(context.Background()),
			memory.New(),
			1*time.Minute,
			30*24*time.Hour,
			nil,
		),
		MockClient: c,
	}
}
