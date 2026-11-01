package claim

import (
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"io"
	"sync"
)

type Connection struct {
	mutex   sync.Mutex
	client  *client.Client
	updates chan Message
	body    io.ReadCloser
}
