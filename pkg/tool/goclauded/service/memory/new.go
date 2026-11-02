package memory

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"time"
)

func New(
	s *store.Store,
	m memoryd.Client,
	g *peer.Guard,
	messages *message.Service,
	clock func() time.Time,
	r library.Reporter,
) *Service {
	return &Service{
		store:          s,
		memory:         m,
		memoryGuard:    g,
		message:        messages,
		clock:          clock,
		reporter:       r,
		lastMemoryPoll: clock().UTC().Format("2006-01-02T15:04:05Z"),
	}
}
