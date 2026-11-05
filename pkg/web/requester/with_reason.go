package requester

import "github.com/funtimecoding/soil/pkg/web/refusal"

func (r *Requester) WithReason(v refusal.Reader) *Requester {
	r.reason = v

	return r
}
