package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestALoadIsARealLink(t *testing.T) {
	assert.String(
		t,
		`<a href="/alfa">Bravo</a>`,
		renderedNode(t, load.New("/alfa", "Bravo").Node()),
	)
}

func TestASwappedLoadGetsItsPathIntoTheTarget(t *testing.T) {
	assert.String(
		t,
		`<a href="/alfa" role="button" class="outline" hx-get="/alfa" hx-target="#charlie" hx-swap="outerHTML">Bravo</a>`,
		renderedNode(
			t,
			load.New("/alfa", "Bravo").
				Swap("#charlie", constant.SwapOuter).
				Button().
				Class("outline").
				Node(),
		),
	)
}
