package transcript

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
)

func New(
	s *store.Store,
	c face.ClaudeSource,
	cache *session_cache.Cache,
	memories *memory.Service,
	index *search_index.Index,
	harbor string,
	sources func() []string,
	exportPath string,
	l *logger.Logger,
	r soilFace.Reporter,
) *Service {
	return &Service{
		store:      s,
		client:     c,
		cache:      cache,
		memory:     memories,
		index:      index,
		harbor:     harbor,
		sources:    sources,
		exportPath: exportPath,
		logger:     l,
		reporter:   r,
	}
}
