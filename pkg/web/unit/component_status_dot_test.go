package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/status_dot"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"testing"
)

func TestStatusDotCarriesItsColourAndTitle(t *testing.T) {
	assert.String(
		t,
		`<span class="status-dot" title="fresh" style="--dot-color: #3dd68c;"></span>`,
		web_tester.New(t).Render(
			status_dot.New("#3dd68c").Title("fresh").Node(),
		),
	)
}

func TestStatusDotNeedsNoTitle(t *testing.T) {
	assert.String(
		t,
		`<span class="status-dot" style="--dot-color: #111;"></span>`,
		web_tester.New(t).Render(status_dot.New("#111").Node()),
	)
}
