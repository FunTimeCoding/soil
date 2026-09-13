package web

import (
	"context"
	"github.com/funtimecoding/soil/pkg/face"
	"net/http"
)

func RecordingMiddleware[F ~func(
	x context.Context,
	w http.ResponseWriter,
	q *http.Request,
	request any,
) (any, error)](
	t face.Recorder,
	skip ...string,
) func(
	F,
	string,
) F {
	skipped := make(map[string]struct{}, len(skip))

	for _, operation := range skip {
		skipped[operation] = struct{}{}
	}

	return func(
		f F,
		operation string,
	) F {
		if _, machine := skipped[operation]; machine {
			return f
		}

		return func(
			x context.Context,
			w http.ResponseWriter,
			q *http.Request,
			request any,
		) (any, error) {
			response, e := f(x, w, q, request)
			RecordTelemetry(t, operation, e)

			return response, e
		}
	}
}
