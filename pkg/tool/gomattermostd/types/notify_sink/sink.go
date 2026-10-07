package notify_sink

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"sync"
)

type Sink struct {
	received []client.NotifyRequest
	mutex    sync.Mutex
}
