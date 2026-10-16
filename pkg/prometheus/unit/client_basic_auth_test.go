package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientSendsBasicAuth(t *testing.T) {
	var seen string
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				user, password, _ := r.BasicAuth()
				seen = join.Empty(user, stringsConstant.Colon, password)
				_, e := w.Write([]byte(`{"status":"success","data":["job"]}`))
				errors.PanicOnError(e)
			},
		),
	)
	defer s.Close()
	c := prometheus.New(
		"127.0.0.1",
		s.Listener.Addr().(*net.TCPAddr).Port,
		false,
		"alfa",
		"bravo",
	)
	defer c.Close()
	labels, e := c.AllLabels()
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"job"}, labels)
	assert.String(t, "alfa:bravo", seen)
}
