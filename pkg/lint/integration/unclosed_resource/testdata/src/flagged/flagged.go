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

func inspect(*Client) {}

func HandedOff() {
	c := newClient()
	inspect(c)
}

func fetch() (*http.Response, error) {
	return http.Get("http://example.com")
}

func UseFetch() int {
	r, _ := fetch()

	return r.StatusCode
}

func follow() (*http.Response, error) {
	response, e := http.Get("http://example.com")

	for range 2 {
		_ = response.Body.Close()
		response, e = http.Get("http://example.com/next")
	}

	return response, e
}

func UseFollow() int {
	r, _ := follow()

	return r.StatusCode
}

func retry() (*http.Response, error) {
	var response *http.Response
	var e error

	for range 2 {
		if response != nil {
			_ = response.Body.Close()
		}

		response, e = http.Get("http://example.com")
	}

	return response, e
}

func UseRetry() int {
	r, _ := retry()

	return r.StatusCode
}

type Factory struct{}

func (f *Factory) Make() *Client {
	return newClient()
}

func UseFactory(f *Factory) {
	made := f.Make()
	made.Work()
}

func (c *Client) With() *Client {
	return c
}

func DropChained() {
	chained := newClient().With()
	chained.Work()
}

type Maker interface {
	Make() *Client
}

func UseLocalMaker() {
	var m Maker = &Factory{}
	local := m.Make()
	local.Work()
}
