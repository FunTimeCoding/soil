package usage_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage_window"
	"github.com/funtimecoding/soil/pkg/web/chart"
	"github.com/funtimecoding/soil/pkg/web/chart/series"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"time"
)

func Chart(window *usage_window.Window) gomponents.Node {
	if window == nil {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("No usage data yet."),
		)
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
