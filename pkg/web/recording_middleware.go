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
) (any, error)](t face.Recorder) func(F, string) F {
	return func(f F, operation string) F {
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
