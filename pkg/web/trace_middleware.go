package web

import (
	"github.com/getsentry/sentry-go"
	"net/http"
)

func TraceMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			q *http.Request,
		) {
			h := sentry.CurrentHub().Clone()
			h.Scope().SetPropagationContext(propagation(q))
			n.ServeHTTP(
				w,
				q.WithContext(sentry.SetHubOnContext(q.Context(), h)),
			)
		},
	)
}
