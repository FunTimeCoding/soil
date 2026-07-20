package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/palette"
	theme "github.com/funtimecoding/soil/pkg/web/theme/constant"
	"github.com/funtimecoding/soil/pkg/web/view"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func New(s *service.Service) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    "Dashboard",
			Path:     "/",
			Category: "navigate",
		},
	)

	return &Server{
		service:  s,
		registry: registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Sprout).
				WithStyle(constant.InlineStyle).
				WithCommandPalette("/palette").
				WithScript("https://cdn.jsdelivr.net/npm/sortablejs@1.15.6/Sortable.min.js").
				WithLiveEndpoint("/event").
				WithFooter(
					html.Script(gomponents.Raw(constant.SortableScript)),
				),
		),
	}
}
