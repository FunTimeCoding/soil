package metric

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"sync"
)

func New(
	port int,
	verbose bool,
	w *sync.WaitGroup,
) *Server {
	var address string

	if port == 0 {
		port = Port
		address = constant.MetricAddress
	} else {
		address = fmt.Sprintf(":%d", port)
	}

	r := prometheus.NewRegistry()
	m := http.NewServeMux()
	m.Handle(
		location.Metrics,
		promhttp.HandlerFor(r, promhttp.HandlerOpts{Registry: r}),
	)

	if w == nil {
		w = NewWaitGroup()
	}

	return &Server{
		verbose:  verbose,
		port:     port,
		context:  context.Background(),
		registry: r,
		server:   web.Server(m, address),
		wait:     w,
	}
}
