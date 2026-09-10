package gauge

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (g *Gauge) Poll() {
	findings, e := g.service.Findings()

	if e != nil {
		g.logger.Structured("finding_gauge_failed", "reason", e.Error())

		return
	}

	counts := map[string]float64{}

	for _, i := range findings {
		counts[i.Kind] += float64(i.Count)
	}

	for _, kind := range constant.FindingKinds {
		g.findings.WithLabelValues(kind).Set(counts[kind])
	}
}
