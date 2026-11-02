package daemon

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (d *Daemon) MountBridge(m *http.ServeMux) {
	m.Handle(webConstant.RootPath, d.client)
}
