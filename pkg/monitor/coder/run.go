package coder

import (
	"context"
	"errors"
	errorLibrary "github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/monitor/coder/server"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() error {
	if len(os.Args) < 2 {
		return errors.New("listen address missing")
	}

	l, listenFail := net.Listen(system.Transmission, os.Args[1])
	errorLibrary.PanicOnError(listenFail)
	log.Printf(
		"listen on %s",
		locator.New(l.Addr().String()).Scheme(constant.Socket).String(),
	)
	s := &http.Server{
		Handler:      server.Server{Logf: log.Printf},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	channelFail := make(chan error, 1)
	go func() {
		log.Println("serve")
		channelFail <- s.Serve(l)
	}()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	log.Println("wait for signal")

	select {
	case selectFail := <-channelFail:
		log.Printf("serve fail: %v", selectFail)
	case s := <-signals:
		log.Printf("terminate: %v", s)
	}

	c, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return s.Shutdown(c)
}
