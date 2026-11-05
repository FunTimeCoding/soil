package claim

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
	"io"
	"sync"
)

type Connection struct {
	mutex   sync.Mutex
	client  face.Client
	updates chan Message
	body    io.ReadCloser
	context context.Context
}
