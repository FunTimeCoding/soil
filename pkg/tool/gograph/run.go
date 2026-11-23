package gograph

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/microsoft/graph"
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/oauth2"
)

func Run() {
	c := graph.NewEnvironment()
	defer c.Close()
	e := c.Login(
		context.Background(),
		func(a *oauth2.DeviceAuthResponse) {
			console.Format(
				"Open %s and enter %s\n",
				a.VerificationURI,
				a.UserCode,
			)
		},
	)

	if validation.Is(e) {
		c.Close()
		system.Exitf(1, "%s\n", e)
	}

	errors.PanicOnError(e)
	console.Format("Signed in; the token is in the Keychain.\n")
}
