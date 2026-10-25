package client

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/route"
)

func (c *Client) Mount(g *guard.Mux) {
	g.WithSession(c.Require)
	g.Open(route.Get(c.signInPath), c.SignIn)
	g.Open(route.Get(constant.CallbackPath), c.Callback)
	g.Open(route.Get(constant.SignOutPath), c.SignOut)
}
