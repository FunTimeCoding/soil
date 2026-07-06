package memory_indexer_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/memory_indexer"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	search http.HandlerFunc,
	list http.HandlerFunc,
) *Tester {
	t.Helper()
	m := http.NewServeMux()
	m.HandleFunc("/api/search", search)
	m.HandleFunc("/api/list", list)
	s := httptest.NewServer(m)
	t.Cleanup(s.Close)
	c, e := client.NewClient(s.URL)
	errors.PanicOnError(e)

	return &Tester{Indexer: memory_indexer.New(c)}
}
