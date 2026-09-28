package server

import (
	"encoding/json"
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (s *Server) register(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body registerRequest

	if e := json.NewDecoder(r.Body).Decode(&body); e != nil {
		http.Error(w, library.InvalidRequestBody, http.StatusBadRequest)

		return
	}

	if len(body.RedirectLocators) == 0 {
		http.Error(w, "redirect_locators required", http.StatusBadRequest)

		return
	}

	grantTypes := body.GrantTypes

	if len(grantTypes) == 0 {
		grantTypes = []string{constant.GrantAuthorizationCode}
	}

	responseTypes := body.ResponseTypes

	if len(responseTypes) == 0 {
		responseTypes = []string{constant.ResponseCode}
	}

	scopes := []string{constant.ScopeOpenIdentity}

	if body.Scope != "" {
		scopes = []string{body.Scope}
	}

	result, e := s.service.RegisterClient(
		service.NewRegisterClientRequest(
			body.RedirectLocators,
			grantTypes,
			responseTypes,
			scopes,
			body.TokenEndpointAuthMethod,
		),
	)

	if e != nil {
		http.Error(w, "registration failed", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
	web.EncodeNotation(
		w,
		&registerResponse{
			ClientIdentifier:        result.ClientIdentifier,
			ClientSecret:            result.ClientSecret,
			RedirectLocators:        result.RedirectLocators,
			GrantTypes:              grantTypes,
			ResponseTypes:           responseTypes,
			TokenEndpointAuthMethod: body.TokenEndpointAuthMethod,
		},
	)
}
