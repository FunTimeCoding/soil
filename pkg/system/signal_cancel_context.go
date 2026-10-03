package system

import (
	"context"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"log"
	"os"
	"os/signal"
)

func SignalCancelContext(l *log.Logger) context.Context {
	channel := make(chan os.Signal, 2)
	signal.Notify(channel)
	result, cancel := context.WithCancel(context.Background())
	go func() {
		for h := range channel {
			if h == constant.RuntimePreemptionSignal {
				continue
			}

			l.Printf("Signal: %+v (%d)\n", h, h)
			cancel()

			return
		}
	}()

	return result
}
