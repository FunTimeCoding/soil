package collect

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/bubbletea/model/monitor/fetch"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system"
	"time"
)

func Loop() {
	fmt.Printf("Commands: %s\n", join.Comma(fetch.List()))
	loopCheck(time.Now())
	ticker := time.NewTicker(5 * time.Minute)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				loopCheck(t)
			}
		}
	}()
	system.KillSignalBlock()
	done <- true
}
