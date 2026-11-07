package unit

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func paletteView() *view.View {
	return view.New(
		layout.New(identity.New("test", "test tool", "test"), "", nil).
			WithItems(navigation_item.New("/alfa", "Alfa")).
			WithCommands(
				palette.NewCommand("Bravo", "/bravo", constant.PaletteAction),
			),
	)
}
