package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/tile"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"maragu.dev/gomponents"
	"testing"
)

func TestATileCarriesItsSearchText(t *testing.T) {
	assert.String(
		t,
		`<div class="tile foxtrot" data-search="golf hotel">india</div>`,
		renderedNode(
			t,
			tile.New(gomponents.Text("india")).
				Terms("golf hotel").
				Class("foxtrot").
				Node(),
		),
	)
}

func TestAGridWithoutSearchIsOnlyTheGrid(t *testing.T) {
	assert.String(
		t,
		`<div class="tile-grid"><div class="tile">india</div></div>`,
		renderedNode(
			t,
			tile_grid.New(
				[]gomponents.Node{tile.New(gomponents.Text("india")).Node()},
			).Node(),
		),
	)
}

func TestASearchableGridFiltersTheGridBelowIt(t *testing.T) {
	page := renderedNode(t, tile_grid.New(nil).Search("Filter…").Node())
	assert.StringContains(
		t,
		`<input type="search" placeholder="Filter…" oninput="filterTiles(this)"><div class="tile-grid"></div><script>`,
		page,
	)
	assert.StringContains(t, "function filterTiles(input)", page)
}
