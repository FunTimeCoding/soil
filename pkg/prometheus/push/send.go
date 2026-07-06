package push

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/push"
)

func Send(p *push.Pusher) {
	errors.PanicOnError(p.Push())
}
