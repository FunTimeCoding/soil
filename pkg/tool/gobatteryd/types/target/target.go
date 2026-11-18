package target

import (
	"github.com/funtimecoding/soil/pkg/nut/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Target struct {
	Server string
	Client face.Client
	Guard  *peer.Guard
}
