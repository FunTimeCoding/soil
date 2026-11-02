package refusal_transport

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/forbidden"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/proxmox/constant"
	"github.com/luthermonson/go-proxmox"
	"net/http"
	"strconv"
	"strings"
)

func (t *Transport) refuse(
	q *http.Request,
	r *http.Response,
) error {
	reason := strings.TrimSpace(
		strings.TrimPrefix(r.Status, strconv.Itoa(r.StatusCode)),
	)
	text := fmt.Sprintf("proxmox %s %s: %s", q.Method, q.URL.Path, reason)

	switch {
	case r.StatusCode == http.StatusBadRequest && fromCaller(q.Context()):
		return validation.New("%s%s", text, parameters(r.Body))
	case r.StatusCode == http.StatusForbidden:
		return fmt.Errorf(
			"%w: %w",
			forbidden.Format("%s", text),
			proxmox.ErrNotAuthorized,
		)
	case r.StatusCode == http.StatusInternalServerError &&
		strings.Contains(reason, constant.ReasonMissing):

		return not_found.Format("%s", text)
	case r.StatusCode == http.StatusInternalServerError &&
		strings.Contains(reason, constant.ReasonLocked):

		return conflict.Format("%s", text)
	}

	return unexpected.Format(
		"%s (status %d)%s",
		text,
		r.StatusCode,
		parameters(r.Body),
	)
}
