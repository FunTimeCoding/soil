package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/dial"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"testing"
)

func TestADialIsATrackWithACenter(t *testing.T) {
	assert.String(
		t,
		`<div class="dial" style="--dial-background: #111"><div class="dial-center"><strong>alfa</strong></div></div>`,
		renderedNode(
			t,
			dial.New("#111").
				Center(html.Strong(gomponents.Text("alfa"))).
				Node(),
		),
	)
}

func TestAFillCoversTheTrackUpToItsShare(t *testing.T) {
	assert.StringContains(
		t,
		`style="--dial-background: conic-gradient(#222 0deg 90.0deg, #111 90.0deg 360deg)"`,
		renderedNode(t, dial.New("#111").Fill("#222", 0.25).Node()),
	)
	assert.StringContains(
		t,
		`style="--dial-background: #222"`,
		renderedNode(t, dial.New("#111").Fill("#222", 1).Node()),
	)
	assert.StringContains(
		t,
		`style="--dial-background: #111"`,
		renderedNode(t, dial.New("#111").Fill("#222", 0).Node()),
	)
}

func TestArcsLayerOverTheTrack(t *testing.T) {
	assert.StringContains(
		t,
		`style="--dial-background: conic-gradient(transparent 0deg 36.0deg, #333 36.0deg 72.0deg, transparent 72.0deg 360deg), #111"`,
		renderedNode(t, dial.New("#111").Arc("#333", 0.1, 0.2).Node()),
	)
}

func TestMarkersTurnFromTheTop(t *testing.T) {
	ring := renderedNode(
		t,
		dial.New("#111").
			Point(0.25, "#444", "bravo").
			Tick(1.5, "#555", "charlie").
			Node(),
	)
	assert.StringContains(
		t,
		`<div class="dial-needle" style="--dial-angle: 0.0deg"><div class="dial-point" title="bravo" style="--dial-color: #444;"></div></div>`,
		ring,
	)
	assert.StringContains(
		t,
		`style="--dial-angle: 270.0deg"><div class="dial-tick" title="charlie"`,
		ring,
	)
}
