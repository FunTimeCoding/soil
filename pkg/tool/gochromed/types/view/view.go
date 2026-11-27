package view

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium/snapshot"
	"github.com/funtimecoding/soil/pkg/browser/chromium/tab"
)

type View struct {
	Nodes  []*snapshot.Node
	Frames []*tab.Tab
}
