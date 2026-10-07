package registry

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (r *Registry) Serve() http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		q *http.Request,
	) {
		query := q.URL.Query().Get("q")
		results := r.Search(query)
		fragment := palette.ResultList(results)
		w.Header().Set(constant.ContentType, constant.MarkupUnicode)
		errors.PanicOnError(fragment.Render(w))
	}
}
