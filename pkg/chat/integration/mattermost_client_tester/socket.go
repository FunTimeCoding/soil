package mattermost_client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/gorilla/websocket"
	"net/http"
)

// The client dials the websocket during construction; hold the
// upgraded connection open until the client side closes it.
func (t *Tester) socket(
	w http.ResponseWriter,
	q *http.Request,
) {
	u := websocket.Upgrader{}
	c, e := u.Upgrade(w, q, nil)

	if e != nil {
		return
	}

	defer errors.LogClose(c)
	t.accept(c)

	for {
		if _, _, f := c.ReadMessage(); f != nil {
			return
		}
	}
}
