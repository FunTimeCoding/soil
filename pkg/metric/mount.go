package metric

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func (m *Metric) Mount(g *guard.Mux) {
	g.OpenMount(webConstant.MetricsPath, m.Exporter())
}
