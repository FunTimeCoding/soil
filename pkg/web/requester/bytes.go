package requester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"io"
)

func (r *Requester) Bytes(q *request.Request) ([]byte, error) {
	response, e := r.Send(q)

	if e != nil {
		return nil, e
	}

	defer errors.LogClose(response.Body)

	return io.ReadAll(response.Body)
}
