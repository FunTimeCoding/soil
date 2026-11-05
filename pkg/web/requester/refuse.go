package requester

import (
	"fmt"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"github.com/funtimecoding/soil/pkg/web/requester/request"
	"net/http"
)

func (r *Requester) refuse(
	o *request.Request,
	q *http.Request,
	response *http.Response,
	body []byte,
) error {
	if o.Refusal != nil {
		if e := o.Refusal(response.StatusCode, body); e != nil {
			return e
		}
	}

	if r.refusal != nil {
		if e := r.refusal(response.StatusCode, body); e != nil {
			return e
		}
	}

	class := r.table.Class(response.StatusCode)

	if c, okay := o.Statuses[response.StatusCode]; okay {
		class = c
	}

	reason := r.reason(body, response.Status)

	if class == errorsConstant.ValidationClass {
		return refusal.Classed(class, reason)
	}

	subject := r.subject

	if subject == "" {
		subject = join.Empty(q.URL.Host, q.URL.Path)
	}

	return refusal.Classed(
		class,
		fmt.Sprintf("%s status: %d: %s", subject, response.StatusCode, reason),
	)
}
