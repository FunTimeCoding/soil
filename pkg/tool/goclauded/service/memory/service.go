package memory

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/message"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	memoryd "github.com/funtimecoding/soil/pkg/tool/gomemoryd/face"
	"time"
)

type Service struct {
	store          *store.Store
	memory         memoryd.Client
	memoryGuard    *peer.Guard
	message        *message.Service
	clock          func() time.Time
	reporter       library.Reporter
	lastMemoryPoll string
}
