package requester

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
)

func (r *Requester) refuse(
	q *http.Request,
	response *http.Response,
	body []byte,
) error {
	if r.refusal != nil {
		if e := r.refusal(response.StatusCode, body); e != nil {
			return e
		}
	}

	if response.StatusCode == http.StatusNotFound {
		return not_found.Format(
			"resource not found: %s%s",
			q.URL.Host,
			q.URL.Path,
		)
	}

	return unexpected.Format(
		"%s%s status: %d: %s",
		q.URL.Host,
		q.URL.Path,
		response.StatusCode,
		reason(body, response.Status),
	)
}
