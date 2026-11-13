package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"testing"
)

func TestPaletteLeadsWithDaemonsWayHome(t *testing.T) {
	c := layout.New(identity.New("alfa", "alfa tool", "alfa"), "", nil).
		WithItems(navigation_item.New("/bravo", "Bravo")).
		Commands()
	assert.Count(t, 2, c)
	assert.String(t, "alfa", c[0].Label)
	assert.String(t, "/", c[0].Path)
	assert.String(t, "Bravo", c[1].Label)
}

func TestLayoutWithNothingToListHasNoWayHome(t *testing.T) {
	assert.Count(
		t,
		0,
		layout.New(identity.New("alfa", "alfa tool", "alfa"), "", nil).Commands(),
	)
}
