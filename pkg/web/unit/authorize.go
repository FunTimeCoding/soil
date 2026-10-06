package unit

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (a *renewing) Authorize(r *http.Request) error {
	web.Bearer(r, a.token)

	return nil
}
