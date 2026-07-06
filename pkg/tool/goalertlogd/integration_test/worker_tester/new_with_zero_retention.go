package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"path/filepath"
	"testing"
	"time"
)

func NewWithZeroRetention(t *testing.T) *Tester {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	t.Cleanup(s.Close)
	c := mock_client.New()
	w := worker.New(
		c,
		s,
		logger.New(context.Background()),
		memory.New(),
		1*time.Minute,
		0,
		nil,
	)

	return &Tester{Store: s, Worker: w, MockClient: c}
}
