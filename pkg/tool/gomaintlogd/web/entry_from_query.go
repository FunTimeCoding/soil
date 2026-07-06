package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"net/http"
)

func (s *Server) entryFromQuery(r *http.Request) *entry.Entry {
	result, f := s.store.Get(
		strings.MustToUnsignedInteger(
			r.URL.Query().Get(constant.Identifier),
		),
	)
	errors.PanicOnError(f)

	return result
}
