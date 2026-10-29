package view

import (
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
)

type View struct {
	Nodes  []*snapshot.Node
	Frames []*tab.Tab
}
