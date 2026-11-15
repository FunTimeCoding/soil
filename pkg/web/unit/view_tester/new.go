package view_tester

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New() *view.View {
	return view.New(
		layout.New(identity.New("test", "test tool", "test"), "", nil),
	)
}
