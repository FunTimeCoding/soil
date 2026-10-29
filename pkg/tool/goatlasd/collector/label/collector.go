package label

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Collector struct {
	netbox face.Client
	guard  *peer.Guard
}
