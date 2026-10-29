package target

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/face"
)

type Target struct {
	Name   string
	Client face.Client
	Guard  *peer.Guard
}
