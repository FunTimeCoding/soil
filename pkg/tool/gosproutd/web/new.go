package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(s *service.Service) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.SessionsTitle,
			constant.SessionsPath,
			web.PaletteNavigate,
		),
	)

	return &Server{
		service:  s,
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeSprout).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.SeedsTitle,
					),
					navigation_item.New(
						constant.SessionsPath,
						constant.SessionsTitle,
					),
					navigation_item.New(
						constant.CountersPath,
						constant.CountersTitle,
					),
				).
				WithScript(
					"https://cdn.jsdelivr.net/npm/sortablejs@1.15.6/Sortable.min.js",
				).
				WithLiveEndpoint(web.LivePath).
				WithFooter(
					html.Script(gomponents.Raw(constant.SortableScript)),
				),
		),
	}
}
