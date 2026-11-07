package option

import "github.com/funtimecoding/soil/pkg/web/option"

type Monitor struct {
	option.Web
	PostgresLocator string
	LitePath        string
}
