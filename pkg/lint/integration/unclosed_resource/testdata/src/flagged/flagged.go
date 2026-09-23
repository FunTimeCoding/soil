package example

import (
	"net/http"
	"os"
)

func DroppedResponse() int {
	r, _ := http.Get("http://example.com")

	return r.StatusCode
}

func DroppedFile() string {
	f, _ := os.Open("example")

	return f.Name()
}

func DroppedClient() {
	c := newClient()
	c.Work()
}
