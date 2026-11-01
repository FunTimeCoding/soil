package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/meter"
	"maragu.dev/gomponents"
	"testing"
)

func TestABareMeterIsANativeProgressBar(t *testing.T) {
	assert.String(
		t,
		`<progress class="meter-bar" value="12" max="40"></progress>`,
		renderedNode(t, meter.New(12, 40).Node()),
	)
}

func TestAColouredMeterSetsPicosProgressVariables(t *testing.T) {
	assert.String(
		t,
		`<progress class="meter-bar" value="0.6" max="1" style="--pico-progress-color: #5b8def; --pico-progress-background-color: #2b2b3a"></progress>`,
		renderedNode(
			t,
			meter.New(0.6, 1).Color("#5b8def").Track("#2b2b3a").Node(),
		),
	)
}

func TestAMeterCarriesItsSizeAsVariables(t *testing.T) {
	assert.String(
		t,
		`<progress class="meter-bar" value="1" max="2" style="--meter-width: 6rem; --meter-height: 0.9rem"></progress>`,
		renderedNode(t, meter.New(1, 2).Width("6rem").Height("0.9rem").Node()),
	)
}

func TestALabelledMeterIsARowOfLabelBarAndValue(t *testing.T) {
	assert.String(
		t,
		`<div class="meter"><span class="meter-label">Load</span><progress class="meter-bar" value="0.42" max="1"></progress><span class="meter-value">0.42</span></div>`,
		renderedNode(t, meter.New(0.42, 1).Label("Load").Shown("0.42").Node()),
	)
}

func TestAGaugeCarriesItsShareColourAndWidthAsVariables(t *testing.T) {
	assert.String(
		t,
		`<div class="meter" data-tooltip="alfa: 7/10"><span class="meter-icon">◆</span><div class="meter-gauge" style="--meter-share: 70.0%; --meter-color: #8a6fb0; --meter-width: 4rem"><div class="meter-gauge-fill"></div></div></div>`,
		renderedNode(
			t,
			meter.New(7, 10).
				Gauge().
				Icon(gomponents.Text("◆")).
				Color("#8a6fb0").
				Width("4rem").
				Title("alfa: 7/10").
				Node(),
		),
	)
}

func TestAGaugePastItsMaximumStopsFull(t *testing.T) {
	assert.StringContains(
		t,
		`--meter-share: 100.0%`,
		renderedNode(t, meter.New(5, 3).Gauge().Node()),
	)
}
