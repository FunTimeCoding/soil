package pool

import "net/http"

func (p *Pool) Transport() *http.Transport {
	return p.transport
}
