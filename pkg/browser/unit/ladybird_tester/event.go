package ladybird_tester

import (
	"github.com/funtimecoding/soil/pkg/browser/types/message"
	"time"
)

func (t *Tester) Event() *message.Reply {
	t.t.Helper()

	select {
	case r := <-t.events:
		return r
	case <-time.After(5 * time.Second):
		t.t.Fatal("no event arrived")

		return nil
	}
}
