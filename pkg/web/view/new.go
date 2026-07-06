package view

import "github.com/funtimecoding/soil/pkg/web/layout"

func New(l *layout.Page) *View {
	return &View{layout: l}
}
