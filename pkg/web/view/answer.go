package view

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"net/http"
)

func (v *View) Answer(
	w http.ResponseWriter,
	q *http.Request,
	r face.Reporter,
	e error,
) {
	v.renderFailure(w, q, caughtStatus(e), layout.ErrorItem(report(r, e)))
}
