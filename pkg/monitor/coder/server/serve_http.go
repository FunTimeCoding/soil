package server

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

func (s Server) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	c, echoFail := websocket.Accept(
		w,
		r,
		&websocket.AcceptOptions{Subprotocols: []string{"echo"}},
	)
	log.Printf("accept: %s\n", r.RemoteAddr)

	if echoFail != nil {
		s.Logf("%v", echoFail)

		return
	}

	defer func() {
		if e := c.CloseNow(); e != nil {
			s.Logf("failed to close connection: %v", e)
		}
	}()
	log.Printf("subprotocol: %v\n", c.Subprotocol())

	if c.Subprotocol() != "echo" {
		errors.LogOnError(
			c.Close(
				websocket.StatusPolicyViolation,
				"client must speak the echo sub-protocol",
			),
		)

		return
	}

	l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)

	for {
		log.Println("wait for echo")
		echoFail = echo(c, l)

		if websocket.CloseStatus(echoFail) == websocket.StatusNormalClosure {
			return
		}

		if echoFail != nil {
			s.Logf("echo fail %v: %v", r.RemoteAddr, echoFail)

			return
		}
	}
}
