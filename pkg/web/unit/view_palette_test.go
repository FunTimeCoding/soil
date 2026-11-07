package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/unit/web_tester"
	"github.com/funtimecoding/soil/pkg/web/view"
	"net/http"
	"testing"
)

func TestThePaletteListsTheItemsThenTheExtras(t *testing.T) {
	r := paletteAnswer(t, paletteView(), true)
	assert.Integer(t, http.StatusOK, r.Code)
	assert.StringContains(t, "Alfa", r.Body.String())
	assert.StringContains(t, "Bravo", r.Body.String())
	assert.StringNotContains(t, "Sign out", r.Body.String())
}

func TestASignedSessionAlsoGetsSignOut(t *testing.T) {
	r := paletteAnswer(t, paletteView(), false)
	assert.StringContains(t, "Sign out", r.Body.String())
}

func TestALayoutWithNothingToListHasNoPalette(t *testing.T) {
	i := identity.New("test", "test tool", "test")
	plain := layout.New(i, "", nil)
	assert.StringNotContains(t, `id="palette"`, web_tester.RenderBrand(plain))
	assert.StringContains(
		t,
		`id="palette"`,
		web_tester.RenderBrand(
			layout.New(i, "", nil).WithItems(
				navigation_item.New("/alfa", "Alfa"),
			),
		),
	)
	assert.StringContains(
		t,
		`id="palette"`,
		web_tester.RenderBrand(
			layout.New(i, "", nil).WithCommands(
				palette.NewCommand("Bravo", "/bravo", constant.PaletteAction),
			),
		),
	)
	assert.Integer(
		t,
		http.StatusNotFound,
		paletteAnswer(t, view.New(plain), true).Code,
	)
}
