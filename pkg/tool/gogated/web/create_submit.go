package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/form"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"net/url"
	"strings"
)

func (s *Server) createSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	rawLocators := r.FormValue("redirect_locators")
	scopes := r.FormValue("scopes")
	var redirectLocators []string

	for _, line := range strings.Split(rawLocators, "\n") {
		trimmed := strings.TrimSpace(line)

		if trimmed != "" {
			redirectLocators = append(redirectLocators, trimmed)
		}
	}

	if len(redirectLocators) == 0 {
		preserve := url.Values{}
		preserve.Set("redirect_locators", rawLocators)
		preserve.Set("scopes", scopes)
		form.Redirect(
			w,
			r,
			constant.CreatePath,
			"At least one redirect locator is required.",
			preserve,
		)

		return
	}

	var scopeSlice []string

	if scopes != "" {
		scopeSlice = strings.Fields(scopes)
	}

	result, e := s.service.RegisterClient(
		service.NewFleetClientRequest(redirectLocators, scopeSlice),
	)
	errors.PanicOnError(e)
	s.view.RenderPage(
		w,
		"Client Created",
		constant.CreatePath,
		html.H1(gomponents.Text("Client Created")),
		createResult(result.ClientIdentifier, result.ClientSecret),
	)
}
