package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/classify"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/failure"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (s *Server) answer(
	x context.Context,
	w http.ResponseWriter,
	e error,
) {
	status := http.StatusInternalServerError

	if a := classify.Answer(e); a != nil {
		status = a.Status
	}

	result := failure.New(classify.Message(e, ""))

	if classify.Reportable(e) {
		result.EventIdentifier = s.reporter.Capture(x, e)
	}

	web.ObjectHeader(w)
	w.WriteHeader(status)
	web.Encode(w, result)
}
