package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/types/register_client"
	"github.com/funtimecoding/soil/pkg/tool/gogated/web/client_page"
	"github.com/funtimecoding/soil/pkg/web/form"
	"net/http"
	"net/url"
	"strings"
)

func (s *Server) createSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	rawLocators := r.FormValue(constant.LocatorsField)
	scopes := r.FormValue(constant.ScopesField)
	var redirectLocators []string

	for _, line := range strings.Split(rawLocators, "\n") {
		trimmed := strings.TrimSpace(line)

		if trimmed != "" {
			redirectLocators = append(redirectLocators, trimmed)
		}
	}

	var scopeSlice []string

	if scopes != "" {
		scopeSlice = strings.Fields(scopes)
	}

	result, e := s.service.RegisterClient(
		register_client.NewFleetRequest(redirectLocators, scopeSlice),
	)

	if validation.Is(e) {
		preserve := url.Values{}
		preserve.Set(constant.LocatorsField, rawLocators)
		preserve.Set(constant.ScopesField, scopes)
		form.Redirect(w, r, constant.CreatePath, e.Error(), preserve)

		return
	}

	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		"Client Created",
		constant.CreatePath,
		client_page.Created(result.ClientIdentifier, result.ClientSecret),
	)
}
