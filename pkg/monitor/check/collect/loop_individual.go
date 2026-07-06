package collect

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"time"
)

func LoopIndividual() {
	fmt.Printf("Commands: %s\n", join.Comma(fetch.List()))
	byName := make(map[string]*LastRun)
	tick(byName, time.Now())
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				tick(byName, t)
			}
		}
	}()
	system.KillSignalBlock()
	done <- true
}
