package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/conversations"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/layout/navigation_item"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	rosterService *roster.Service,
	sessionService *session.Service,
	labelService *label.Service,
	messageService *message.Service,
	eventService *event.Service,
	transcriptService *transcript.Service,
	usageService *usage.Service,
	searchService *search.Service,
	n face.EventNotifier,
	f face.Reporter,
) *Server {
	r := registry.New()
	r.Register(
		palette.NewCommand(
			constant.DashboardTitle,
			constant.DashboardPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.SessionsTitle,
			constant.SessionsPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.MessagesTitle,
			constant.MessagesPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.HistoryTitle,
			constant.HistoryPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.CoverageTitle,
			constant.CoveragePath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.UsageTitle,
			constant.UsagePath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.StatusTitle,
			constant.StatusPath,
			webConstant.PaletteNavigate,
		),
		palette.NewCommand(
			constant.ConversationsTitle,
			constant.ConversationsPath,
			webConstant.PaletteNavigate,
		),
	)

	return &Server{
		roster:     rosterService,
		session:    sessionService,
		label:      labelService,
		message:    messageService,
		event:      eventService,
		transcript: transcriptService,
		usage:      usageService,
		reporter:   f,
		notifier:   n,
		conversations: conversations.New(
			rosterService,
			searchService,
			transcriptService,
		),
		registry: r,
		view: view.New(
			layout.New(constant.Identity).
				WithTheme(webConstant.ThemeHearth).
				WithStyle(
					join.Empty(
						webConstant.ComponentStyle,
						constant.InlineStyle,
						webConstant.ChartStyle,
					),
				).
				WithCommandPalette(webConstant.PalettePath).
				WithLiveEndpoint(webConstant.LivePath).
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
