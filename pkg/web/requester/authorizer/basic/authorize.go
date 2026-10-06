package basic

import "net/http"

func (b *Basic) Authorize(r *http.Request) error {
	r.SetBasicAuth(b.user, b.password)

	return nil
}
