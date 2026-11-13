package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/meter"
	"maragu.dev/gomponents"
	"testing"
)

func TestBareMeterIsNativeProgressBar(t *testing.T) {
	assert.String(
		t,
		`<progress class="meter-bar" value="12" max="40"></progress>`,
		renderedNode(t, meter.New(12, 40).Node()),
	)
}

func TestColouredMeterSetsPicosProgressVariables(t *testing.T) {
	assert.String(
		t,
		`<progress class="meter-bar" value="0.6" max="1" style="--pico-progress-color: #5b8def; --pico-progress-background-color: #2b2b3a"></progress>`,
		renderedNode(
			t,
			meter.New(0.6, 1).Color("#5b8def").Track("#2b2b3a").Node(),
		),
	)
}

func TestMeterCarriesItsSizeAsVariables(t *testing.T) {
	assert.String(
		t,
		`<progress class="meter-bar" value="1" max="2" style="--meter-width: 6rem; --meter-height: 0.9rem"></progress>`,
		renderedNode(t, meter.New(1, 2).Width("6rem").Height("0.9rem").Node()),
	)
}

func TestLabelledMeterIsRowOfLabelBarAndValue(t *testing.T) {
	assert.String(
		t,
		`<div class="meter"><span class="meter-label">Load</span><progress class="meter-bar" value="0.42" max="1"></progress><span class="meter-value">0.42</span></div>`,
		renderedNode(t, meter.New(0.42, 1).Label("Load").Shown("0.42").Node()),
	)
}

func TestTallMeterWithIconKeepsNativeBar(t *testing.T) {
	assert.String(
		t,
		`<div class="meter" data-tooltip="alfa: 7/10"><span class="meter-icon">◆</span><progress class="meter-bar" value="7" max="10" style="--pico-progress-color: #8a6fb0; --meter-width: 4rem; --meter-height: 1.125rem"></progress></div>`,
		renderedNode(
			t,
			meter.New(7, 10).
				Icon(gomponents.Text("◆")).
				Color("#8a6fb0").
				Width("4rem").
				Height("1.125rem").
				Title("alfa: 7/10").
				Node(),
		),
	)
}
