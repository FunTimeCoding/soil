package requester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/requester/face"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"io"
	"net/http"
	"time"
)

func (r *Requester) Send(q *request.Request) (*http.Response, error) {
	renewed := false

	for attempt := 0; ; attempt++ {
		h, e := r.build(q)

		if e != nil {
			return nil, e
		}

		response, e := r.client.Do(h)
		retry := q.Idempotent() && attempt < constant.Retries

		if e != nil {
			if retry && healable(e) {
				time.Sleep(r.delay(attempt, nil))

				continue
			}

			return nil, transportFailure(h, e)
		}

		if succeeded(response.StatusCode) {
			return response, nil
		}

		body, f := io.ReadAll(
			io.LimitReader(response.Body, constant.MaximumRefusal),
		)

		if f != nil {
			body = nil
		}

		errors.LogClose(response.Body)
		renewable, canRenew := r.authorizer.(face.Renewable)

		if response.StatusCode == http.StatusUnauthorized && canRenew &&
			!renewed {
			if f := renewable.Renew(); f != nil {
				return nil, f
			}

			renewed = true

			continue
		}

		if retry && transient(response.StatusCode) {
			time.Sleep(r.delay(attempt, response))

			continue
		}

		return nil, r.refuse(h, response, body)
	}
}
