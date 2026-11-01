package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/option"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	memoryMock "github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	queryMock "github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	directory := t.TempDir()
	l := logger.New(t.Context())
	s := store_tester.NewOnFile(t, filepath.Join(directory, "goclauded.sqlite"))
	c := mock_client.New()
	mc := memoryMock.New()
	t.Cleanup(mc.Close)
	q := queryMock.New()
	t.Cleanup(q.Close)
	n := mock_notifier.New()
	r := memory.New()
	harbor := t.TempDir()
	c.Harbor = harbor
	o := option.New()
	o.LitePath = filepath.Join(directory, "goclauded.sqlite")
	o.SessionExportPath = t.TempDir()
	o.Harbor = harbor
	o.Clock = s.Clock()
	o.Claude = func() face.ClaudeSource { return c }
	o.Memory = func() memoryd.Client { return mc }
	o.Query = func() queryd.Client { return q }
	o.Notifier = func() face.Notifier { return n }
	v := service.New(o, l, r)
	t.Cleanup(v.Close)
	x := connection.New(l, filepath.Join(directory, constant.SearchIndexFile))
	t.Cleanup(func() { errors.PanicClose(x) })

	return &Tester{
		Store:    s,
		t:        t,
		Service:  v,
		Harbor:   harbor,
		Client:   c,
		Query:    q,
		Notifier: n,
		Memory:   mc,
		Search:   search_index.New(x),
		Reporter: r,
	}
}
