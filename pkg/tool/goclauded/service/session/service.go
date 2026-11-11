package session

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
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

type Service struct {
	store      *store.Store
	label      *label.Service
	message    *message.Service
	cruise     *cruise.Service
	transcript *transcript.Service
	client     face.ClaudeSource
	cache      *session_cache.Cache
	index      *search_index.Index
	notifier   soilFace.EventNotifier
	clock      func() time.Time
	harbor     string
	sources    func() []string
}
