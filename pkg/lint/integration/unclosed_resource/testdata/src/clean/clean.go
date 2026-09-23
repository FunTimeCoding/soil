package example

import (
	"example/errors"
	"net/http"
	"os"
)

type holder struct {
	response *http.Response
}

func consume(r *http.Response) {}

func ClosedDirect() int {
	r, _ := http.Get("http://example.com")
	defer r.Body.Close()

	return r.StatusCode
}

func ClosedHelper() int {
	r, _ := http.Get("http://example.com")
	defer errors.PanicClose(r.Body)

	return r.StatusCode
}

func ClosedLog() string {
	f, _ := os.Open("example")
	defer errors.LogClose(f)

	return f.Name()
}

func Returned() *http.Response {
	r, _ := http.Get("http://example.com")

	return r
}

func PassedOn() {
	r, _ := http.Get("http://example.com")
	consume(r)
}

func Stored() *holder {
	r, _ := http.Get("http://example.com")

	return &holder{response: r}
}

func Captured() func() int {
	r, _ := http.Get("http://example.com")

	return func() int { return r.StatusCode }
}
