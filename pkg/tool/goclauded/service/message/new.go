package message

import (
	library "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/stopper"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store"
	"time"
)

func New(
	s *store.Store,
	n face.Notifier,
	stop *stopper.Stopper,
	clock func() time.Time,
	r library.Reporter,
) *Service {
	return &Service{
		store:    s,
		notifier: n,
		stop:     stop,
		clock:    clock,
		reporter: r,
	}
}
