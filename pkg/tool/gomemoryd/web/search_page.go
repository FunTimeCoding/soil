package web

import (
	library "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) searchPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	query := r.URL.Query().Get(constant.Query)

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, s.searchResults(query))

		return
	}

	s.view.RenderPage(
		w,
		constant.SearchTitle,
		web.SearchPath,
		html.H3(gomponents.Text(constant.SearchTitle)),
		html.Form(
			gomponents.Attr(web.FormMethod, http.MethodGet),
			gomponents.Attr(web.FormAction, web.SearchPath),
			html.Input(
				html.ID(constant.SearchControlMark),
				html.Type("search"),
				html.Name(constant.Query),
				gomponents.Attr("placeholder", constant.SearchPlaceholder),
				gomponents.Attr("value", query),
				gomponents.Attr("autocomplete", "off"),
				extended.Get(web.SearchPath),
				extended.Trigger(web.TriggerType),
				extended.Target(
					join.Empty(library.Hash, constant.SearchResultsMark),
				),
				extended.Swap(web.SwapInner),
				extended.Indicator(
					join.Empty(library.Hash, constant.SearchIndicatorMark),
				),
			),
			html.Button(
				html.Type(web.FormSubmit),
				gomponents.Text(constant.SearchTitle),
			),
		),
		layout.Indicator(
			constant.SearchIndicatorMark,
			constant.SearchIndicatorText,
		),
		html.Div(html.ID(constant.SearchResultsMark), s.searchResults(query)),
	)
}
