package startup

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/roster"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/search"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
)

func New(
	rosters *roster.Service,
	transcripts *transcript.Service,
	searches *search.Service,
) *Service {
	return &Service{
		roster:     rosters,
		transcript: transcripts,
		search:     searches,
	}
}
