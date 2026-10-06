package bearer

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
)

func (b *Bearer) Authorize(r *http.Request) error {
	web.Bearer(r, b.token)

	return nil
}
