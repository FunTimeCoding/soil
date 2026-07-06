package wdutil

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	library "github.com/funtimecoding/soil/pkg/time"
	"time"
)

func Check() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go func() {
		past := NotAvailable

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				r := collect()

				if past != r.Sequence {
					fmt.Printf(
						"%s change %s\n",
						t.Format(library.DateSecond),
						r.Sequence,
					)
				}

				past = r.Sequence
			}
		}
	}()
	system.KillSignalBlock()
	done <- true
}
