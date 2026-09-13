package view

import (
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (v *View) RenderStatus(
	w http.ResponseWriter,
	message string,
	kind string,
) {
	v.RenderFragment(w, layout.StatusItem(message, kind))
}
