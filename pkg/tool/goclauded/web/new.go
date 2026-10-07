package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversations"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
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
		palette.NewCommand(
			constant.MessagesTitle,
			constant.MessagesPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.HistoryTitle,
			constant.HistoryPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CoverageTitle,
			constant.CoveragePath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.UsageTitle,
			constant.UsagePath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.StatusTitle,
			constant.StatusPath,
			web.PaletteNavigate,
		),
		palette.NewCommand(
			constant.ConversationsTitle,
			constant.ConversationsPath,
			web.PaletteNavigate,
		),
	)

	return &Server{
		service:       s,
		notifier:      s.Notifier(),
		conversations: conversations.New(s),
		registry:      r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(web.ThemeHearth).
				WithStyle(join.Empty(constant.InlineStyle, web.ChartStyle)).
				WithCommandPalette(web.PalettePath).
				WithLiveEndpoint(web.LivePath).
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
					navigation_item.New(
						constant.CoveragePath,
						constant.CoverageTitle,
					),
					navigation_item.New(
						constant.UsagePath,
						constant.UsageTitle,
					),
					navigation_item.New(
						constant.StatusPath,
						constant.StatusTitle,
					),
					navigation_item.NewExternal(
						constant.ConversationsPath,
						constant.ConversationsTitle,
					),
				),
		),
	}
}
