package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/go-library/pkg/web/authorization/client"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"github.com/funtimecoding/go-library/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	theme "github.com/funtimecoding/go-library/pkg/web/theme/constant"
	"github.com/funtimecoding/go-library/pkg/web/view"
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
			Category: navigateCategory,
		},
		palette.Command{
			Label:    constant.HeatmapTitle,
			Path:     constant.HeatmapPath,
			Category: navigateCategory,
		},
	)
	labels := map[string]bool{}

	for _, entry := range b.Entries() {
		labels[entry.Label] = true
		registry.Register(
			palette.Command{
				Label:    entry.Label,
				Path:     entry.Link,
				Category: linkCategory,
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
				WithStyle(inlineCSS).
				WithCommandPalette(palettePath).
				WithLiveEndpoint(eventPath).
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
					html.Script(gomponents.Raw(beaconJS)),
				),
		),
	}
}
