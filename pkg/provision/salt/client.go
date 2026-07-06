package salt

import "github.com/funtimecoding/soil/pkg/provision/salt/basic"

type Client struct {
	basic          *basic.Client
	authentication string
}
