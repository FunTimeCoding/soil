package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Flight struct {
	option.Web
	PostgresLocator string
	LitePath        string
	Predicate       string
	Retention       time.Duration
	Clock           func() time.Time
	Command         func(s ...string) (string, error)
	Stream          func(s ...string) face.Stream
}
