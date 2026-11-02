package conversations

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
)

type Server struct {
	roster     *roster.Service
	search     *search.Service
	transcript *transcript.Service
}
