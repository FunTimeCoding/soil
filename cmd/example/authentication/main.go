package main

import (
	"github.com/funtimecoding/soil/pkg/text/multi_line"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/authenticator"
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/funtimecoding/soil/pkg/web/request_context"
	"net/http"
)

func main() {
	a := authenticator.New()
	m := http.NewServeMux()
	m.HandleFunc(
		location.Root,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			c.SetLastLocation()
			l := multi_line.New()

			if a.LoggedIn(c) {
				l.Add("logged in")
			} else {
				l.Add("not logged in")
			}

			c.WriteOkay(l.Render())
		},
	)
	m.HandleFunc(
		location.Status,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			c.SetLastLocation()
			l := multi_line.New()
			l.Format("Session: %s", a.Session(c))
			c.WriteOkay(l.Render())
		},
	)
	m.HandleFunc(
		location.Login,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			a.AddressLogin(c)
			c.Redirect(c.LastLocation())
		},
	)
	m.HandleFunc(
		location.Logout,
		func(
			w http.ResponseWriter,
			e *http.Request,
		) {
			c := request_context.New(w, e)
			a.Logout(c)
			c.Redirect(c.LastLocation())
		},
	)
	web.Listen(m)
}
