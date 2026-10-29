package view

import (
	"maragu.dev/gomponents"
	"net/http"
)

func (v *View) RenderPage(
	w http.ResponseWriter,
	title string,
	path string,
	content ...gomponents.Node,
) {
	v.RenderPageWithStatus(w, http.StatusOK, title, path, content...)
}
