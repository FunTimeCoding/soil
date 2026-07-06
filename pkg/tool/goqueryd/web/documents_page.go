package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) documentsPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.PathValue("name")
	sourceType := r.URL.Query().Get(constant.SourceType)
	page := 1

	if v := r.URL.Query().Get("page"); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			page = n
		}
	}

	limit := 20
	offset := (page - 1) * limit
	var metadata map[string]string

	if sourceType != "" {
		metadata = map[string]string{constant.SourceType: sourceType}
	}

	outcome, e := s.service.ListDocuments(
		name,
		metadata,
		limit+1,
		offset,
		false,
	)

	if e != nil {
		s.view.RenderPage(
			w,
			name,
			constant.CollectionsPath,
			html.P(
				gomponents.Textf(
					"Failed to load documents: %s",
					e.Error(),
				),
			),
		)

		return
	}

	results := outcome.Results
	hasMore := len(results) > limit

	if hasMore {
		results = results[:limit]
	}

	table := documentsTable(results)
	navigation := documentsNavigation(name, page, hasMore, sourceType)
	isHTMX := r.Header.Get("HX-Request") == "true"

	if isHTMX {
		errors.PanicOnError(
			gomponents.Group(
				[]gomponents.Node{table, navigation},
			).Render(w),
		)

		return
	}

	var filters []gomponents.Node
	var sourceTypeFacet *store.Facet

	for i, f := range outcome.Facets {
		if f.Key == constant.SourceType && f.Values != nil {
			sourceTypeFacet = &outcome.Facets[i]

			break
		}
	}

	if sourceTypeFacet != nil && len(sourceTypeFacet.Values) > 1 {
		clearAttrs := []gomponents.Node{
			html.Type("radio"),
			html.Name(constant.SourceType),
			html.Value(""),
			gomponents.Attr(
				"hx-get",
				fmt.Sprintf("/documents/%s", name),
			),
			gomponents.Attr("hx-include", "#documents-filters"),
			gomponents.Attr("hx-target", "#documents-content"),
			gomponents.Attr("hx-swap", "innerHTML"),
		}

		if sourceType == "" {
			clearAttrs = append(
				clearAttrs,
				gomponents.Attr("checked", ""),
			)
		}

		filters = append(
			filters,
			html.Label(
				html.Input(clearAttrs...),
				gomponents.Text("All"),
			),
		)

		for value := range sourceTypeFacet.Values {
			attrs := []gomponents.Node{
				html.Type("radio"),
				html.Name(constant.SourceType),
				html.Value(value),
				gomponents.Attr(
					"hx-get",
					fmt.Sprintf("/documents/%s", name),
				),
				gomponents.Attr("hx-include", "#documents-filters"),
				gomponents.Attr("hx-target", "#documents-content"),
				gomponents.Attr("hx-swap", "innerHTML"),
			}

			if sourceType == value {
				attrs = append(attrs, gomponents.Attr("checked", ""))
			}

			filters = append(
				filters,
				html.Label(
					html.Input(attrs...),
					gomponents.Textf(
						"%s (%d)",
						value,
						sourceTypeFacet.Values[value],
					),
				),
			)
		}
	}

	var content []gomponents.Node
	content = append(
		content,
		html.H3(gomponents.Text(name)),
	)

	if len(filters) > 0 {
		content = append(
			content,
			html.Form(
				html.ID("documents-filters"),
				gomponents.Group(filters),
			),
		)
	}

	content = append(
		content,
		html.Div(
			html.ID("documents-content"),
			table,
			navigation,
		),
	)
	s.view.RenderPage(w, name, constant.CollectionsPath, content...)
}
