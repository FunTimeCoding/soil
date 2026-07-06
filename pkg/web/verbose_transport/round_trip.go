package verbose_transport

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (d *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	return web.VerboseTrip(d.base, r)
}
