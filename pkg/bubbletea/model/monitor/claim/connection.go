package claim

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/client"
	"io"
	"sync"
)

type Connection struct {
	mutex   sync.Mutex
	client  *client.Client
	updates chan Message
	body    io.ReadCloser
	context context.Context
}
