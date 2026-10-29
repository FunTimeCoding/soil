package lease

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/face"
)

type Collector struct {
	opnsense face.Client
	guard    *peer.Guard
}
