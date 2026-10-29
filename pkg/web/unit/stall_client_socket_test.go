package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestStallClientSocketDialsThePath(t *testing.T) {
	directory, e := os.MkdirTemp("/tmp", "socket")
	assert.FatalOnError(t, e)
	t.Cleanup(func() { assert.FatalOnError(t, os.RemoveAll(directory)) })
	path := filepath.Join(directory, "alfa.sock")
	l, f := net.Listen("unix", path)
	assert.FatalOnError(t, f)
	s := httptest.NewUnstartedServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				_, g := w.Write([]byte("bravo"))
				assert.FatalOnError(t, g)
			},
		),
	)
	s.Listener = l
	s.Start()
	defer s.Close()
	p := stall_client.New().Socket(path).Build()
	defer p.Close()
	assert.String(t, "bravo", web.GetString(p.Client(), "http://goprocessd/"))
}
