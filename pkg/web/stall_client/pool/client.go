package pool

import "net/http"

func (p *Pool) Client() *http.Client {
	return p.client
}
