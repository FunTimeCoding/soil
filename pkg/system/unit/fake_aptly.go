package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/linux/debian/aptly"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func fakeAptly(
	t *testing.T,
	status int,
	body string,
	seen *aptlySeen,
) *aptly.Client {
	t.Helper()
	s := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				b, e := io.ReadAll(r.Body)
				errors.PanicOnError(e)
				*seen = aptlySeen{
					Method: r.Method,
					Path:   r.URL.Path,
					Query:  r.URL.Query(),
					Body:   string(b),
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(status)
				_, f := w.Write([]byte(body))
				errors.PanicOnError(f)
			},
		),
	)
	t.Cleanup(s.Close)
	u, e := url.Parse(s.URL)

	if e != nil {
		t.Fatal(e)
	}

	host, port, f := net.SplitHostPort(u.Host)

	if f != nil {
		t.Fatal(f)
	}

	number, g := strconv.Atoi(port)

	if g != nil {
		t.Fatal(g)
	}

	result := aptly.New(host, number, true, "", "")
	t.Cleanup(result.Close)

	return result
}
