package header

import "net/http"

func (h *Header) Authorize(r *http.Request) error {
	r.Header.Set(h.name, h.value)

	return nil
}
