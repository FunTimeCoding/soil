package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"time"
)

type Flight struct {
	Address         string
	PostgresLocator string
	LitePath        string
	Predicate       string
	ServiceTokens   []string
	Retention       time.Duration
	Clock           func() time.Time
	Command         func(s ...string) (string, error)
	Stream          func(s ...string) face.Stream
}
