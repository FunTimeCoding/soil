package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration/tester"
	"sync"
	"testing"
)

func TestStatusDuringReloadProcfile(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	stop := make(chan struct{})
	var group sync.WaitGroup
	group.Add(1)
	go func() {
		defer group.Done()

		for {
			select {
			case <-stop:
				return
			default:
				s.Send("status")
			}
		}
	}()

	for i := range 20 {
		if i%2 == 0 {
			s.WriteProcfile(
				"alfa: sleep 60\nbravo: sleep 60\ncharlie: sleep 60\n",
			)
		} else {
			s.WriteProcfile("alfa: sleep 60\nbravo: sleep 60\n")
		}

		s.Send("reload-procfile")
	}

	close(stop)
	group.Wait()
}
