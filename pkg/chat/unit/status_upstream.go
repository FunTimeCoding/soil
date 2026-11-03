package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/mattermost/mattermost/server/public/model"
	"net/http"
)

func statusUpstream(m *http.ServeMux) {
	answer := func(
		path string,
		status int,
	) {
		m.HandleFunc(
			path,
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				w.WriteHeader(status)
				web.Encode(
					w,
					&model.AppError{
						Id:         "refused",
						Message:    http.StatusText(status),
						StatusCode: status,
					},
				)
			},
		)
	}
	answer("/api/v4/posts/forbidden", http.StatusForbidden)
	answer("/api/v4/posts/missing", http.StatusNotFound)
	answer("/api/v4/posts/unauthorized", http.StatusUnauthorized)
	answer("/api/v4/posts/teapot", http.StatusTeapot)
	m.HandleFunc(
		"/api/v4/posts/proxied",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			w.Header().Set(webConstant.ContentType, webConstant.MarkupUnicode)
			w.WriteHeader(http.StatusServiceUnavailable)
			_, e := w.Write([]byte("<html>maintenance</html>"))
			errors.PanicOnError(e)
		},
	)
	m.HandleFunc(
		"/api/v4/posts/hollow",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			w.WriteHeader(http.StatusNotFound)
		},
	)
	m.HandleFunc(
		"/api/v4/posts/truncated",
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			w.Header().Set(webConstant.ContentType, webConstant.Object)
			w.Header().Set("Content-Length", "64")
			w.WriteHeader(http.StatusOK)
			_, e := w.Write([]byte("{"))
			errors.PanicOnError(e)
		},
	)
}
