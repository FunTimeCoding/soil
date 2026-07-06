package authenticator

import "github.com/funtimecoding/soil/pkg/web/request_context"

func (a *Authenticator) LoggedIn(c *request_context.Context) bool {
	return a.Session(c) != ""
}
