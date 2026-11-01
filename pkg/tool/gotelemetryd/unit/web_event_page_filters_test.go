package unit

import (
	"bytes"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/types/query_option"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web/event_page"
	"testing"
)

func TestRemovingOneFilterKeepsTheOthers(t *testing.T) {
	o := query_option.New()
	o.Tool = "alfa"
	o.Kind = "bravo"
	var b bytes.Buffer
	assert.FatalOnError(t, event_page.Filters(o).Render(&b))
	assert.StringContains(
		t,
		`<a href="/events?kind=bravo&amp;page=1" class="chip">tool: alfa ×</a>`,
		b.String(),
	)
	assert.StringContains(
		t,
		`<a href="/events?page=1&amp;tool=alfa" class="chip">kind: bravo ×</a>`,
		b.String(),
	)
}
