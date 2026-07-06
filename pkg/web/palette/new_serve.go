package palette

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func NewServe(registry *Registry) http.HandlerFunc {
	return func(
		w http.ResponseWriter,
		q *http.Request,
	) {
		query := q.URL.Query().Get("q")
		results := registry.Search(query)
		fragment := resultList(results)
		w.Header().Set(constant.ContentType, constant.MarkupUnicode)
		errors.PanicOnError(fragment.Render(w))
	}
}
