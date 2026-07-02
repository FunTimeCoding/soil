package base

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/go-library/pkg/generative/model_context_server"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/mock_client"
	generated "github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/model_context"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/server"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/model_context/mock_recorder"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			"critical",
			"Memory above 90%",
		),
	)
	c.Add(
		alert.NewBasic(
			"fp2",
			"DiskFull",
			"warning",
			"Disk usage above 85%",
		),
	)
	l := logger.New(context.Background())
	r := memory.New()
	w := worker.New(c, s, l, r, 1*time.Minute, 30*24*time.Hour, nil)
	w.Poll()
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			generated.HandlerFromMux(
				generated.NewStrictHandler(server.New(s, w, r), nil),
				m,
			)
			model_context.New(
				s,
				w,
				r,
				mock_recorder.New(),
				constant.DefaultVersion,
			).Mount(m)
			web.New(s, w).Mount(m)
		},
	)

	return &Server{
		Store:         s,
		Worker:        w,
		MockClient:    c,
		ContextServer: v,
	}
}
