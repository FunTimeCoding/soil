package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/chart"
	"github.com/funtimecoding/soil/pkg/web/chart/series"
	"strings"
	"testing"
	"time"
)

func renderChart(c *chart.Chart) string {
	var b strings.Builder
	errors.PanicOnError(c.Render().Render(&b))

	return b.String()
}

func TestChartStepPath(t *testing.T) {
	start := time.Date(2026, 8, 26, 21, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 7)
	weekly := series.New("Weekly", "chart-series-a").
		Add(start.Add(24*time.Hour), 10).
		Add(start.Add(48*time.Hour), 20)
	markup := renderChart(
		chart.New(start, end).
			WithNow(start.Add(72 * time.Hour)).
			WithSeries(weekly).
			WithGuide().
			WithProjection(),
	)
	assert.StringContains(
		t,
		`d="M 132.6 212.0 H 221.1 V 190.0 H 309.7"`,
		markup,
	)
	assert.StringContains(t, "chart-guide", markup)
	assert.StringContains(t, "chart-now", markup)
	assert.StringContains(t, ">47%<", markup)
}

func TestChartProjectionCapped(t *testing.T) {
	start := time.Date(2026, 8, 26, 21, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 7)
	weekly := series.New("Weekly", "chart-series-a").
		Add(start.Add(24*time.Hour), 80)
	markup := renderChart(
		chart.New(start, end).
			WithNow(start.Add(24 * time.Hour)).
			WithSeries(weekly).
			WithProjection(),
	)
	assert.StringContains(t, ">100%<", markup)
}
