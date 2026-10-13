package bearer

import "net/http"

type Transport struct {
	Base http.RoundTripper
}
