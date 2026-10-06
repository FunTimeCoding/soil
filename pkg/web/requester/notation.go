package requester

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
)

func (r *Requester) Notation(
	q *request.Request,
	out any,
) error {
	response, e := r.Send(q)

	if e != nil {
		return e
	}

	defer errors.LogClose(response.Body)

	if f := json.NewDecoder(response.Body).Decode(out); f != nil {
		return unexpected.Format(
			"%s%s: answer is not JSON: %s",
			response.Request.URL.Host,
			response.Request.URL.Path,
			f,
		)
	}

	return nil
}
