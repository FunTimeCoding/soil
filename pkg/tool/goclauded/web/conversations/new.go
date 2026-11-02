package conversations

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
)

func New(
	rosterService *roster.Service,
	searchService *search.Service,
	transcriptService *transcript.Service,
) *Server {
	return &Server{
		roster:     rosterService,
		search:     searchService,
		transcript: transcriptService,
	}
}
