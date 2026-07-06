package example_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/monitor/gorilla/helper"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/location"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

func Run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	u := url.URL{
		Scheme: constant.Socket,
		Host: web.AddressHostPort(
			constant.Localhost,
			constant.ListenPort,
		),
		Path: location.Echo,
	}
	log.Printf("connect to %s", u.String())
	l := helper.Dial(u)
	defer errors.LogClose(l)
	done := make(chan struct{})
	go func() {
		defer close(done)

		for {
			_, message, e := l.ReadMessage()

			if e != nil {
				log.Printf("read: %s\n", e)

				return
			}

			log.Printf("receive: %s", message)
		}
	}()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			if e := l.WriteMessage(
				websocket.TextMessage,
				[]byte(t.String()),
			); e != nil {
				log.Printf("write: %s\n", e)

				return
			}
		case <-interrupt:
			log.Println("interrupt")

			if e := l.WriteMessage(
				websocket.CloseMessage,
				websocket.FormatCloseMessage(
					websocket.CloseNormalClosure,
					"",
				),
			); e != nil {
				log.Printf("write close: %s\n", e)

				return
			}

			select {
			case <-done:
			case <-time.After(time.Second):
			}

			return
		}
	}
}
