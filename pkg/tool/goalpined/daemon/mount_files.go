package daemon

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
)

func (d *Daemon) MountFiles(g *guard.Mux) {
	g.OpenMount(stringsConstant.Slash, http.FileServer(http.Dir(d.packageRoot)))
}
