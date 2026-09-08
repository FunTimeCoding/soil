package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"sync"
)

type notifySink struct {
	received []client.NotifyRequest
	mutex    sync.Mutex
}
