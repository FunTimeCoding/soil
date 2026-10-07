package renewing_authorizer

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (a *Authorizer) Authorize(r *http.Request) error {
	web.Bearer(r, a.Token)

	return nil
}
