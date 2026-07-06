package main

import (
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/system"
	"log"
)

func main() {
	m := metric.New(0, true, nil)
	log.Println("starting")
	m.Start()
	log.Println("started")
	system.KillSignalBlock()
	log.Println("stopping")
	m.Stop()
	log.Println("stopped")
}
