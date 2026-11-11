package event

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/document"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/memory"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

type Service struct {
	store    *store.Store
	message  *message.Service
	memory   *memory.Service
	document *document.Service
	notifier face.EventNotifier
	clock    func() time.Time
	reporter face.Reporter
}
