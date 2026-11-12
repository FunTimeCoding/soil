package claim

import (
	"context"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/claim/message"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/face"
	"io"
	"sync"
)

type Connection struct {
	mutex   sync.Mutex
	client  face.Client
	updates chan message.Message
	body    io.ReadCloser
	context context.Context
}
