package option

import (
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Telemetry struct {
	option.Web
	PostgresLocator string
	LitePath        string
	Clock           func() time.Time
}
