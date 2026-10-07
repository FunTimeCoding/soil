package result

import "net/http"

type Authorize struct {
	Code                 string
	CodeVerifier         string
	RedirectLocator      string
	AuthenticationCookie *http.Cookie
}
