package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *store.Store,
	p *worker.Worker,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.RecentTitle,
			Path:     constant.RecentPath,
			Category: web.PaletteNavigate,
		},
	)

	return &Server{
		store:    s,
		worker:   p,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeSentinel).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(web.PalettePath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.RecentPath,
						constant.RecentTitle,
					),
				),
		),
	}
}
