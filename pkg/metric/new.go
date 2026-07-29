package metric

import (
	"context"
	"github.com/funtimecoding/soil/pkg/metric/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
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
	if port == 0 {
		port = environment.FallbackInteger(
			constant.PortEnvironment,
			constant.Port,
		)
	}

	// Metrics listen on all interfaces - Prometheus scrapes over the network
	address := web.AddressHostPort("", port)
	r := prometheus.NewRegistry()
	m := http.NewServeMux()
	m.Handle(
		webConstant.LocationMetrics,
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
