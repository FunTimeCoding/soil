package kubernetes

import (
	"github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Collector struct {
	kubernetes face.Client
	guard      *peer.Guard
}
