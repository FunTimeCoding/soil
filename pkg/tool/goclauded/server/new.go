package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/cruise"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"time"
)

func New(
	rosterService *roster.Service,
	sessionService *session.Service,
	labelService *label.Service,
	messageService *message.Service,
	cruiseService *cruise.Service,
	eventService *event.Service,
	transcriptService *transcript.Service,
	usageService *usage.Service,
	searchService *search.Service,
	l *logger.Logger,
	r face.Reporter,
	hold time.Duration,
) *Server {
	return &Server{
		roster:     rosterService,
		session:    sessionService,
		label:      labelService,
		message:    messageService,
		cruise:     cruiseService,
		event:      eventService,
		transcript: transcriptService,
		usage:      usageService,
		search:     searchService,
		logger:     l,
		reporter:   r,
		hold:       hold,
	}
}
