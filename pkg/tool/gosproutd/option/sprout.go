package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Sprout struct {
	option.Web
	LitePath      string
	SeedDirectory string
	Clock         func() time.Time
	Notifier      func() face.EventNotifier
}
