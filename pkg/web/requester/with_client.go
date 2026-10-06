package requester

import "net/http"

func (r *Requester) WithClient(c *http.Client) *Requester {
	r.client = c

	return r
}
