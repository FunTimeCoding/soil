package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/cruise"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/usage"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	roster     *roster.Service
	session    *session.Service
	label      *label.Service
	message    *message.Service
	cruise     *cruise.Service
	event      *event.Service
	transcript *transcript.Service
	usage      *usage.Service
	search     *search.Service
	reporter   face.Reporter
	notifier   face.EventNotifier
	view       *view.View
}
