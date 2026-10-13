package relay

import (
	"context"
	"github.com/coder/websocket"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goffrelay/constant"
	"github.com/funtimecoding/soil/pkg/tool/goffrelay/relay/option"
	"log"
	"os"
)

func Run(
	o *option.Relay,
	caller string,
) {
	if caller != "" && caller != o.Extension {
		errors.Printf("firefox relay: refused caller %s\n", caller)
		os.Exit(1)
	}

	x, cancel := context.WithCancel(context.Background())
	defer cancel()
	c, _, e := websocket.Dial(x, join.Empty(constant.Scheme, o.Host), nil)

	if e != nil {
		errors.Printf("firefox relay: dial %s: %v\n", o.Host, e)
		os.Exit(1)
	}

	defer func() { errors.LogOnError(c.CloseNow()) }()
	c.SetReadLimit(constant.FrameLimit)
	log.Printf("firefox relay: connected to %s", o.Host)
	closed := make(chan struct{}, 2)
	go forwardReply(x, c, closed)
	go forwardRequest(x, c, closed)
	<-closed
}
