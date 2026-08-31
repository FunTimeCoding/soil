package web

import (
	"github.com/funtimecoding/soil/pkg/web/chart"
	"github.com/funtimecoding/soil/pkg/web/chart/series"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func (s *Server) usageChart() gomponents.Node {
	window := s.service.UsageWindow()

	if window == nil {
		return html.P(gomponents.Text("No usage data yet."))
	}

	weekly := series.New("Weekly", "chart-series-a")

	for _, snapshot := range window.Rate {
		weekly.Add(snapshot.CreatedAt, float64(snapshot.SevenDayPercent))
	}

	fable := series.New("Fable", "chart-series-b")

	for _, snapshot := range window.Fable {
		fable.Add(snapshot.CreatedAt, float64(snapshot.Percent))
	}

	return chart.New(window.Start, window.End).
		WithNow(time.Now()).
		WithSeries(weekly, fable).
		WithGuide().
		WithProjection().
		Render()
}
