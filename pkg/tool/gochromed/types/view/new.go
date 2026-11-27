package view

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/snapshot"
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
)

func New(
	nodes []*snapshot.Node,
	frames []*tab.Tab,
) *View {
	return &View{Nodes: nodes, Frames: frames}
}
