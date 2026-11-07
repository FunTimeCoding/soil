package view

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (v *View) Favicon(
	w http.ResponseWriter,
	q *http.Request,
) {
	favicon := v.layout.Favicon()

	if len(favicon) == 0 {
		http.NotFound(w, q)

		return
	}

	w.Header().Set(constant.ContentType, constant.Graphic)
	web.WriteBytes(w, http.StatusOK, favicon)
}
