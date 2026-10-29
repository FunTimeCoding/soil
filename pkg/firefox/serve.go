package firefox

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
	"log"
	"net/http"
)

func (c *Client) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	connection, e := websocket.Accept(
		w,
		r,
		&websocket.AcceptOptions{InsecureSkipVerify: true},
	)

	if e != nil {
		log.Printf("firefox bridge: accept: %v", e)

		return
	}

	connection.SetReadLimit(constant.ReadLimit)
	c.mutex.Lock()

	if c.closed {
		c.mutex.Unlock()
		errors.LogOnError(connection.CloseNow())

		return
	}

	c.connection = connection
	c.mutex.Unlock()
	log.Printf("firefox bridge: extension connected")

	defer func() {
		c.mutex.Lock()

		if c.connection == connection {
			c.connection = nil
		}

		closed := c.closed
		c.mutex.Unlock()

		if !closed {
			errors.LogOnError(connection.CloseNow())
		}

		log.Printf("firefox bridge: extension disconnected")
	}()
	c.readLoop(connection)
}
