package callback

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/funtimecoding/soil/pkg/web/request_context"
	"net/http"
)

func New(
	port int,
	verbose bool,
) *Server {
	if port == 0 {
		port = constant.ListenPort
	}

	m := http.NewServeMux()
	result := &Server{
		verbose:    verbose,
		port:       port,
		context:    context.Background(),
		server:     web.ServerPort(m, port),
		callbackCh: make(chan string, 1),
		errorCh:    make(chan error, 1),
	}
	m.Handle(
		location.Callback,
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				c := request_context.New(w, r)
				code := c.Query("code")

				if code == "" {
					result.errorCh <- fmt.Errorf(
						"no authorization code received",
					)

					return
				}

				c.WriteOkay("Callback received")
				result.callbackCh <- code
			},
		),
	)

	return result
}
