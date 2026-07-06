package firefox

import (
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/errors"
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

	log.Printf("firefox bridge: extension connected")
	c.mutex.Lock()
	c.connection = connection
	c.mutex.Unlock()

	defer func() {
		c.mutex.Lock()
		c.connection = nil
		c.mutex.Unlock()
		errors.LogOnError(connection.CloseNow())
		log.Printf("firefox bridge: extension disconnected")
	}()
	c.readLoop(connection)
}
