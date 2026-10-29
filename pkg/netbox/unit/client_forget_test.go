package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestAForgottenClientReadsTheRacksAgain(t *testing.T) {
	var requests atomic.Int32
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				_ *http.Request,
			) {
				requests.Add(1)
				w.Header().Set("Content-Type", "application/json")
				_, e := w.Write(
					[]byte(
						`{"count":1,"results":[{"id":1,"url":"/rack/1","display":"alfa","name":"alfa","site":{"id":1,"url":"/site/1","display":"bravo","name":"bravo","slug":"bravo"}}]}`,
					),
				)
				errors.PanicOnError(e)
			},
		),
	)
	t.Cleanup(s.Close)
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	t.Cleanup(c.Close)
	c.MustRacks()
	c.MustRacks()
	assert.Integer(t, 1, requests.Load())
	c.Forget()
	c.MustRacks()
	assert.Integer(t, 2, requests.Load())
}
