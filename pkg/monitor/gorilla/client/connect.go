package client

import (
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/helper"
	"log"
)

func (c *Client) Connect() {
	log.Printf("connect to %s", c.locator.Host)
	c.connection = helper.Dial(c.locator)
	go func() {
		defer close(c.done)

		for {
			_, text, e := c.connection.ReadMessage()

			if e != nil {
				log.Printf("read: %s\n", e)

				return
			}

			log.Printf("receive: %s", text)
			c.receive = append(c.receive, string(text))
		}
	}()
}
