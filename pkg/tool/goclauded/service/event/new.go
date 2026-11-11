package event

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/document"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	messages *message.Service,
	memories *memory.Service,
	documents *document.Service,
	n face.EventNotifier,
	clock func() time.Time,
	r face.Reporter,
) *Service {
	return &Service{
		store:    s,
		message:  messages,
		memory:   memories,
		document: documents,
		notifier: n,
		clock:    clock,
		reporter: r,
	}
}
