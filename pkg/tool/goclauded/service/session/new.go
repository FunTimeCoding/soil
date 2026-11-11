package session

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/cruise"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/transcript"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	labels *label.Service,
	messages *message.Service,
	cruises *cruise.Service,
	transcripts *transcript.Service,
	c face.ClaudeSource,
	cache *session_cache.Cache,
	index *search_index.Index,
	n face.Notifier,
	clock func() time.Time,
	harbor string,
	sources func() []string,
) *Service {
	return &Service{
		store:      s,
		label:      labels,
		message:    messages,
		cruise:     cruises,
		transcript: transcripts,
		client:     c,
		cache:      cache,
		index:      index,
		notifier:   n,
		clock:      clock,
		harbor:     harbor,
		sources:    sources,
	}
}
