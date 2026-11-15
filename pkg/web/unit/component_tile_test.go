package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/component/tile"
	"github.com/funtimecoding/soil/pkg/web/component/tile_grid"
	"github.com/funtimecoding/soil/pkg/web/web_tester"
	"maragu.dev/gomponents"
	"testing"
)

func TestTileCarriesItsSearchText(t *testing.T) {
	assert.String(
		t,
		`<div class="tile foxtrot" data-search="golf hotel">india</div>`,
		web_tester.Render(
			t,
			tile.New(gomponents.Text("india")).
				Terms("golf hotel").
				Class("foxtrot").
				Node(),
		),
	)
}

func TestGridWithoutSearchIsOnlyGrid(t *testing.T) {
	assert.String(
		t,
		`<div class="tile-grid"><div class="tile">india</div></div>`,
		web_tester.Render(
			t,
			tile_grid.New(
				[]gomponents.Node{tile.New(gomponents.Text("india")).Node()},
			).Node(),
		),
	)
}

func TestSearchableGridFiltersGridBelowIt(t *testing.T) {
	page := web_tester.Render(t, tile_grid.New(nil).Search("Filter…").Node())
	assert.StringContains(
		t,
		`<input type="search" placeholder="Filter…" oninput="filterTiles(this)"><div class="tile-grid"></div><script>`,
		page,
	)
	assert.StringContains(t, "function filterTiles(input)", page)
}
