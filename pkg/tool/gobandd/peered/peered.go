package peered

import (
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Peered struct {
	client face.BandClient
	guard  *peer.Guard
}
