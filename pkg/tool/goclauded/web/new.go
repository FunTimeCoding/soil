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
		view: view.New(
			layout.New(constant.Identity, webConstant.ThemeHearth, favicon).
				WithStyle(
					join.Empty(constant.InlineStyle, webConstant.ChartStyle),
				).
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
						constant.PermissionsPath,
						constant.PermissionsTitle,
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
