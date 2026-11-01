package service_tester

import (
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	q := mock_client.New()
	t.Cleanup(q.Close)
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "memory.db")
	o.Query = func() face.Client { return q }
	result := service.New(o, logger.New(t.Context()))
	t.Cleanup(result.Close)

	return &Tester{t: t, Service: result, Indexer: q}
}
