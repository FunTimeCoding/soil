package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(c face.NetboxSource) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.BookmarkTitle,
			Path:     constant.BookmarkPath,
			Category: web.PaletteNavigate,
		},
	)

	return &Server{
		client:   c,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeAmethyst).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(
						constant.BookmarkPath,
						constant.BookmarkTitle,
					),
				),
		),
	}
}
