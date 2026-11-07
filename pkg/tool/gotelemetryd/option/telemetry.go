package option

import "github.com/funtimecoding/soil/pkg/web/option"

type Telemetry struct {
	option.Web
	PostgresLocator string
	LitePath        string
}
