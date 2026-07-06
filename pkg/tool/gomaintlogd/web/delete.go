package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		s.store.Delete(
			strings.MustToUnsignedInteger(
				r.URL.Query().Get(constant.Identifier),
			),
		),
	)
	w.Header().Set(web.ContentType, web.MarkupUnicode)
	w.WriteHeader(http.StatusOK)
}
