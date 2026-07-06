package push

import (
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/prometheus/client_golang/prometheus/push"
)

func New(
	host string,
	port int,
	secure bool,
	job string,
) *push.Pusher {
	return push.New(web.Link(host, port, secure), job)
}
