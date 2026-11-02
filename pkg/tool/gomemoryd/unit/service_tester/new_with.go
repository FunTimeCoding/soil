package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"path/filepath"
	"testing"
)

func newWith(
	t *testing.T,
	hiddenTag string,
	root string,
) *Tester {
	t.Helper()
	l := logger.New(t.Context())
	r := memory.New()
	d := connection.New(l, filepath.Join(t.TempDir(), "memory.db"))
	t.Cleanup(func() { errors.PanicClose(d) })
	q := mock_client.New()
	t.Cleanup(q.Close)
	g := peer.New(constant.Identity.Name(), q.Address(), q, l, r)
	t.Cleanup(g.Close)

	return &Tester{
		t:        t,
		Service:  service.New(store.New(d), q, g, hiddenTag, root, r),
		Indexer:  q,
		Reporter: r,
	}
}
