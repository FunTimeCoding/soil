package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"net/http"
)

func (s *Server) paletteMemories(
	w http.ResponseWriter,
	_ *http.Request,
) {
	fragment := palette.SearchFragment(
		"/palette/memories/search",
		"Search memories...",
	)
	w.Header().Set(webConstant.ContentType, webConstant.MarkupUnicode)
	errors.PanicOnError(fragment.Render(w))
}
