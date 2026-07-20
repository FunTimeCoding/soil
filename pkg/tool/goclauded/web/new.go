package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversations"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
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
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.SessionsTitle,
			Path:     constant.SessionsPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.MessagesTitle,
			Path:     constant.MessagesPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.HistoryTitle,
			Path:     constant.HistoryPath,
			Category: "navigate",
		},
		palette.Command{
			Label:    constant.ConversationsTitle,
			Path:     constant.ConversationsPath,
			Category: "navigate",
		},
	)

	return &Server{
		service:       s,
		notifier:      s.Notifier(),
		conversations: conversations.New(s),
		registry:      registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(theme.Hearth).
				WithStyle(constant.InlineStyle).
				WithCommandPalette("/palette").
				WithLiveEndpoint("/event").
				WithBrandNode(
					html.Strong(
						html.Span(
							html.ID("sse-dot"),
							html.Class(
								"status-dot status-disconnected",
							),
						),
						gomponents.Text(constant.Identity.Name()),
					),
				).
				WithItems(
					navigation_item.New(
						constant.DashboardPath,
						constant.DashboardTitle,
					),
					navigation_item.New(
						constant.SessionsPath,
						constant.SessionsTitle,
					),
					navigation_item.New(
						constant.MessagesPath,
						constant.MessagesTitle,
					),
					navigation_item.New(
						constant.HistoryPath,
						constant.HistoryTitle,
					),
					navigation_item.NewExternal(
						constant.ConversationsPath,
						constant.ConversationsTitle,
					),
				).
				WithFooter(
					html.Script(
						gomponents.Raw(constant.ConnectionIndicatorScript),
					),
				),
		),
	}
}
