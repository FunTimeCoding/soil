package notifier

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
)

func New(
	c face.Client,
	source string,
	g *peer.Guard,
) *Notifier {
	return &Notifier{connector: c, source: source, guard: g}
}
