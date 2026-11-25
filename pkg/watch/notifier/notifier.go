package notifier

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

type Notifier struct {
	connector face.Client
	source    string
	guard     *peer.Guard
}
