package tester

import "net/http"

type AuthorizeResult struct {
	Code                 string
	CodeVerifier         string
	RedirectLocator      string
	AuthenticationCookie *http.Cookie
}
