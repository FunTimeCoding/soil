package refusal

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/reported"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"net/http"
)

func New(
	operation string,
	r *http.Response,
	body []byte,
) error {
	reason := Reason(body, r.Status)

	if event := Event(body); event != "" {
		return reported.New(operation, reason, event)
	}

	if e := NewDaemonTable().Classify(
		r.StatusCode,
		fmt.Sprintf("%s: %s", operation, reason),
	); e != nil {
		return e
	}

	return unexpected.Format(
		"%s status: %d: %s",
		operation,
		r.StatusCode,
		reason,
	)
}
