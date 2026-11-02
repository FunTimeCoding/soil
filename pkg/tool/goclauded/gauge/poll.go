package gauge

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func (g *Gauge) Poll() {
	findings, e := g.findings()

	if e != nil {
		if classify.Reportable(e) {
			g.reporter.CaptureException(e)
		}

		return
	}

	counts := map[string]float64{}

	for _, i := range findings {
		counts[i.Kind] += float64(i.Count)
	}

	for _, kind := range constant.FindingKinds {
		g.gauge.WithLabelValues(kind).Set(counts[kind])
	}
}
