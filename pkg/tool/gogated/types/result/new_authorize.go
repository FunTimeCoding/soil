package result

import "net/http"

func NewAuthorize(
	code string,
	codeVerifier string,
	redirectLocator string,
	authenticationCookie *http.Cookie,
) *Authorize {
	return &Authorize{
		Code:                 code,
		CodeVerifier:         codeVerifier,
		RedirectLocator:      redirectLocator,
		AuthenticationCookie: authenticationCookie,
	}
}
