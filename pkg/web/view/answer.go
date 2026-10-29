package view

import (
	"github.com/funtimecoding/soil/pkg/errors/classify"
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
	status := http.StatusInternalServerError

	if a := classify.Answer(e); a != nil {
		status = a.Status
	}

	v.renderFailure(w, q, status, layout.ErrorItem(report(r, e)))
}
