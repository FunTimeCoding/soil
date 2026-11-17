package target

import (
	"github.com/funtimecoding/soil/pkg/patrol/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Target struct {
	Instance string
	Client   face.PatrolClient
	Guard    *peer.Guard
}
