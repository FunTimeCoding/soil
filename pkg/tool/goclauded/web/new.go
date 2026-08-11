package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversations"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *service.Service) *Server {
	registry := palette.NewRegistry()
	registry.Register(
		palette.Command{
			Label:    constant.DashboardTitle,
			Path:     constant.DashboardPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.SessionsTitle,
			Path:     constant.SessionsPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.MessagesTitle,
			Path:     constant.MessagesPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.HistoryTitle,
			Path:     constant.HistoryPath,
			Category: web.PaletteNavigate,
		},
		palette.Command{
			Label:    constant.ConversationsTitle,
			Path:     constant.ConversationsPath,
			Category: web.PaletteNavigate,
		},
	)

	return &Server{
		service:       s,
		notifier:      s.Notifier(),
		conversations: conversations.New(s),
		registry:      registry,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeHearth).
				WithStyle(constant.InlineStyle).
				WithCommandPalette("/palette").
				WithLiveEndpoint("/event").
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
				),
		),
	}
}
