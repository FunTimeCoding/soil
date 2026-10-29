package option

import (
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/face"
	"time"
)

func New() *Sprout {
	return &Sprout{
		Clock:    time.Now,
		Notifier: func() face.EventNotifier { return notifier.New() },
	}
}
