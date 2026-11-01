package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
)

func New() *Memory {
	return &Memory{Query: func() face.Client { return client.NewEnvironment() }}
}
