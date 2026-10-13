package pool

import "net/http"

func (p *Pool) Halting() *http.Client {
	return p.halting
}
