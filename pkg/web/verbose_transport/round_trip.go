package verbose_transport

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
	"net/http"
	"net/http/httputil"
)

func (d *Transport) RoundTrip(r *http.Request) (*http.Response, error) {
	request, requestFail := httputil.DumpRequestOut(r, true)
	errors.PanicOnError(requestFail)
	log.Printf("REQUEST:\n%s\n", string(request))
	s, e := d.base.RoundTrip(r)

	if s != nil {
		response, responseFail := httputil.DumpResponse(s, true)
		errors.PanicOnError(responseFail)
		log.Printf("RESPONSE:\n%s\n", string(response))
	}

	return s, e
}
