package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"time"
)

type Sprout struct {
	Address       string
	ServiceTokens []string
	LitePath      string
	SeedDirectory string
	Clock         func() time.Time
	Notifier      func() face.EventNotifier
}
