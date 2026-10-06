package face

import "net/http"

type Authorizer interface {
	Authorize(r *http.Request) error
}
