package memory_indexer_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func UnavailableFixture() http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		_ *http.Request,
	) {
		w.Header().Set(constant.ContentType, constant.Object)
		w.WriteHeader(http.StatusServiceUnavailable)
		_, e := w.Write(
			[]byte(`{"error":"embedder unreachable since 14:02: no route"}`),
		)
		errors.PanicOnError(e)
	}
}
