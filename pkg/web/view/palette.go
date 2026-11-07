package view

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"net/http"
)

func (v *View) Palette(
	w http.ResponseWriter,
	q *http.Request,
) {
	if !v.layout.HasCommandPalette() {
		http.NotFound(w, q)

		return
	}

	r := registry.New()
	r.Register(v.layout.Commands()...)

	if guard.Subject(q) != "" {
		r.Register(
			palette.NewCommand(
				constant.SignOutTitle,
				constant.SignOutPath,
				constant.PaletteAction,
			),
		)
	}

	r.Serve()(w, q)
}
