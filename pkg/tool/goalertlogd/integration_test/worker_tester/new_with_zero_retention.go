package worker_tester

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
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
