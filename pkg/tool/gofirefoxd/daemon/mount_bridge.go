package daemon

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (d *Daemon) MountBridge(m *http.ServeMux) {
	m.Handle(constant.RootPath, d.client)
}
