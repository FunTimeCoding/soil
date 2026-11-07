package option

import "github.com/funtimecoding/soil/pkg/web/authorization/client"

type Web struct {
	Address       string
	ServiceTokens []string
	Authorization func() *client.Client
}
