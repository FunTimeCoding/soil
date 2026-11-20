package firefox_tester

import (
	"github.com/funtimecoding/soil/pkg/firefox"
	"net/http/httptest"
)

func (o *Tester) Bridge() (*firefox.Client, *httptest.Server) {
	o.t.Helper()
	c := firefox.New()
	s := httptest.NewServer(c)
	o.t.Cleanup(s.Close)
	o.t.Cleanup(c.Close)

	return c, s
}
