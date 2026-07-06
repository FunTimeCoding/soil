package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
	"net/http"
	"net/http/httputil"
)

func VerboseTrip(
	t http.RoundTripper,
	q *http.Request,
) (*http.Response, error) {
	request, requestFail := httputil.DumpRequestOut(q, true)
	errors.PanicOnError(requestFail)
	log.Printf("REQUEST:\n%s\n", string(request))
	s, e := t.RoundTrip(q)

	if s != nil {
		response, responseFail := httputil.DumpResponse(s, true)
		errors.PanicOnError(responseFail)
		log.Printf("RESPONSE:\n%s\n", string(response))
	}

	return s, e
}
