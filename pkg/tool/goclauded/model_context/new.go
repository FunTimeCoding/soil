package model_context

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/mark/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
)

func New(
	rosterService *roster.Service,
	labelService *label.Service,
	messageService *message.Service,
	eventService *event.Service,
	memoryService *memory.Service,
	usageService *usage.Service,
	searchService *search.Service,
	r face.Reporter,
	l *logger.Logger,
	t face.Recorder,
) *Server {
	result := &Server{
		server:    server.New(constant.Identity, r).WithRecorder(t).Server(),
		roster:    rosterService,
		label:     labelService,
		message:   messageService,
		event:     eventService,
		memory:    memoryService,
		usage:     usageService,
		search:    searchService,
		reporter:  r,
		logger:    l,
		telemetry: t,
	}
	result.register()
	result.registerResources()

	return result
}
