package transcript

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

type Service struct {
	store      *store.Store
	client     face.ClaudeSource
	cache      *session_cache.Cache
	memory     *memory.Service
	index      *search_index.Index
	harbor     string
	sources    func() []string
	exportPath string
	clock      func() time.Time
	logger     *logger.Logger
	reporter   soilFace.Reporter
}
