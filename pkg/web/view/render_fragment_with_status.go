package view

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderFragmentWithStatus(
	w http.ResponseWriter,
	fragment gomponents.Node,
	message string,
	kind string,
) {
	v.RenderFragment(
		w,
		gomponents.Group(
			[]gomponents.Node{fragment, layout.StatusItem(message, kind)},
		),
	)
}
