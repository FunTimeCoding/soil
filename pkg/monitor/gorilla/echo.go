package gorilla

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/monitor/constant"
	"log"
	"net/http"
)

func echo(
	w http.ResponseWriter,
	r *http.Request,
) {
	c, upgradeFail := constant.Upgrader.Upgrade(w, r, nil)

	if upgradeFail != nil {
		log.Printf("upgrade: %s\n", upgradeFail)

		return
	}

	defer errors.LogClose(c)

	for {
		messageType, message, readFail := c.ReadMessage()

		if readFail != nil {
			log.Printf("read: %s\n", readFail)

			break
		}

		log.Printf("receive: %s\n", message)

		if e := c.WriteMessage(messageType, message); e != nil {
			log.Printf("write: %s\n", e)

			break
		}
	}
}
