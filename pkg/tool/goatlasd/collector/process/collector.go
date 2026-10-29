package process

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

type Collector struct {
	process face.Client
	guard   *peer.Guard
	place   string
}
