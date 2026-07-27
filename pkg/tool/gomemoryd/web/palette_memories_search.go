package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) paletteMemoriesSearch(
	w http.ResponseWriter,
	q *http.Request,
) {
	query := q.URL.Query().Get("q")
	var items []palette.SearchItem

	if query != "" {
		results, e := s.service.SearchMemories(
			fmt.Sprintf("%s*", query),
			10,
			"",
			"",
			constant.AllScope,
		)
		errors.PanicOnError(e)

		for _, m := range results {
			items = append(
				items,
				palette.SearchItem{
					Label:       m.Name,
					Description: m.Description,
					Path:        fmt.Sprintf("/memories/%d", m.Identifier),
					Category:    m.Type,
				},
			)
		}
	}

	fragment := palette.SearchResultList(items)
	w.Header().Set(webConstant.ContentType, webConstant.MarkupUnicode)
	errors.PanicOnError(fragment.Render(w))
}
