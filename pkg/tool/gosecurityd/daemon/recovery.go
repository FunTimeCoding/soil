package daemon

import (
	"github.com/funtimecoding/soil/pkg/face"
	"net/http"
)

func (d *Daemon) Recovery(r face.Reporter) func(http.Handler) http.Handler {
	return d.web.Recovery(r)
}
