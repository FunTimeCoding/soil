package router

import (
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/router/client"
	"net"
)

func (r *Router) FindClient(i net.IP) *client.Client {
	for _, l := range r.Client {
		if l.Address.Equal(i) {
			return l
		}
	}

	return nil
}
