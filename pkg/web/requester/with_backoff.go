package requester

import "time"

func (r *Requester) WithBackoff(d time.Duration) *Requester {
	r.backoff = d

	return r
}
