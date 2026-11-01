package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/status_dot"
	"testing"
)

func TestAStatusDotCarriesItsColourAndTitle(t *testing.T) {
	assert.String(
		t,
		`<span class="status-dot" title="fresh" style="--dot-color: #3dd68c;"></span>`,
		renderedNode(t, status_dot.New("#3dd68c").Title("fresh").Node()),
	)
}

func TestAStatusDotNeedsNoTitle(t *testing.T) {
	assert.String(
		t,
		`<span class="status-dot" style="--dot-color: #111;"></span>`,
		renderedNode(t, status_dot.New("#111").Node()),
	)
}
