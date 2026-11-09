package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"time"
)

func New() *Memory {
	return &Memory{
		Clock: time.Now,
		Query: func() face.Client { return client.NewEnvironment() },
	}
}
