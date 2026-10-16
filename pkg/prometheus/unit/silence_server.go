package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/http/httptest"
)

func silenceServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				_, e := w.Write([]byte(constant.FixtureSilences))
				errors.PanicOnError(e)
			},
		),
	)
}
