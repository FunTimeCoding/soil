package client

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/google/uuid"
	"net/http"
	"net/url"
)

func (c *Client) SignIn(
	w http.ResponseWriter,
	r *http.Request,
) {
	verifier := generateVerifier()
	challenge := computeChallenge(verifier)
	state := uuid.New().String()
	returnPath := r.URL.Query().Get("return")

	if returnPath == "" {
		returnPath = "/"
	}

	b, e := json.Marshal(
		&FlowState{
			Verifier:        verifier,
			State:           state,
			ReturnPath:      returnPath,
			CallbackLocator: c.callbackLocator,
		},
	)
	errors.PanicOnError(e)
	encrypted, e := c.encrypt(b)
	errors.PanicOnError(e)
	http.SetCookie(
		w,
		&http.Cookie{
			Name:     flowCookie,
			Value:    encrypted,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
			SameSite: http.SameSiteLaxMode,
		},
	)
	authorizeLocator := fmt.Sprintf(
		"%s/authorize?response_type=code&client_id=%s&redirect_uri=%s&scope=%s&state=%s&code_challenge=%s&code_challenge_method=S256",
		c.issuer,
		url.QueryEscape(c.identifier),
		url.QueryEscape(c.callbackLocator),
		url.QueryEscape(defaultScope),
		url.QueryEscape(state),
		url.QueryEscape(challenge),
	)
	http.Redirect(w, r, authorizeLocator, http.StatusFound)
}
