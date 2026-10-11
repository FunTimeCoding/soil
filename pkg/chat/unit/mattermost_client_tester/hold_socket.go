package mattermost_client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/gorilla/websocket"
	"net/http"
)

func (t *Tester) holdSocket(
	w http.ResponseWriter,
	q *http.Request,
) {
	if t.refused() {
		w.WriteHeader(http.StatusServiceUnavailable)

		return
	}

	u := websocket.Upgrader{}
	c, e := u.Upgrade(w, q, nil)

	if e != nil {
		return
	}

	defer errors.LogClose(c)
	t.accept(c)
	inner := c.PingHandler()
	c.SetPingHandler(
		func(s string) error {
			t.countPing()

			return inner(s)
		},
	)

	for {
		if _, _, f := c.ReadMessage(); f != nil {
			return
		}
	}
}
