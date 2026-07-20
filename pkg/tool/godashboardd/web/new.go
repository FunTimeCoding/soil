package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	theme "github.com/funtimecoding/soil/pkg/web/theme/constant"
	"github.com/funtimecoding/soil/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(
	b *board.Board,
	v *service.Service,
	c *store.Store,
	authorization *client.Client,
) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: constant.NavigateCategory,
		},
		palette.Command{
			Label:    constant.HeatmapTitle,
			Path:     constant.HeatmapPath,
			Category: constant.NavigateCategory,
		},
	)
	labels := map[string]bool{}

	for _, entry := range b.Entries() {
		labels[entry.Label] = true
		registry.Register(
			palette.Command{
				Label:    entry.Label,
				Path:     entry.Link,
				Category: constant.LinkCategory,
			},
		)
	}

	return &Server{
		board:         b,
		service:       v,
		store:         c,
		authorization: authorization,
		labels:        labels,
		registry:      registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Sentinel).
				WithStyle(constant.InlineStyle).
				WithCommandPalette(constant.PalettePath).
				WithLiveEndpoint(constant.EventPath).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.HeatmapPath,
						constant.HeatmapTitle,
					),
				).
				WithFooter(
					html.Script(gomponents.Raw(constant.BeaconScript)),
				),
		),
	}
}
