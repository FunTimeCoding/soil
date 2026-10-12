package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(c face.NetboxSource) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.BookmarkTitle,
			constant.BookmarkPath,
			webConstant.PaletteNavigate,
		),
	)

	return &Server{
		client:   c,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeAmethyst).
				WithCommandPalette(webConstant.PalettePath).
				WithItems(
					navigation_item.New(
						constant.BookmarkPath,
						constant.BookmarkTitle,
					),
				),
		),
	}
}
