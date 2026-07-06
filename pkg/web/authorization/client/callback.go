package client

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (c *Client) Callback(
	w http.ResponseWriter,
	r *http.Request,
) {
	cookie, e := r.Cookie(flowCookie)

	if e != nil {
		http.Error(w, "missing flow state", http.StatusBadRequest)

		return
	}

	b, e := c.decrypt(cookie.Value)

	if e != nil {
		http.Error(w, "invalid flow state", http.StatusBadRequest)

		return
	}

	var flow FlowState
	errors.PanicOnError(json.Unmarshal(b, &flow))
	state := r.URL.Query().Get("state")

	if state != flow.State {
		http.Error(w, "state mismatch", http.StatusBadRequest)

		return
	}

	code := r.URL.Query().Get("code")

	if code == "" {
		http.Error(w, "missing authorization code", http.StatusBadRequest)

		return
	}

	tokens := c.exchangeCode(code, flow.Verifier, flow.CallbackLocator)
	identifierToken := tokens.IdentifierToken

	if identifierToken == "" {
		http.Error(
			w,
			"missing identifier token",
			http.StatusInternalServerError,
		)

		return
	}

	c.ensureProvider()
	verified, e := c.verifier.Verify(context.Background(), identifierToken)

	if e != nil {
		http.Error(w, "token verification failed", http.StatusUnauthorized)

		return
	}

	encrypted, e := c.encrypt([]byte(verified.Subject))
	errors.PanicOnError(e)
	web.SetCookie(w, subjectCookie, encrypted)
	http.SetCookie(
		w,
		&http.Cookie{
			Name:   flowCookie,
			MaxAge: -1,
			Path:   "/",
		},
	)
	http.Redirect(w, r, flow.ReturnPath, http.StatusFound)
}
