package basic

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/session"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
)

type Client struct {
	base      *locator.Locator
	session   *session.Session
	requester *requester.Requester
}
