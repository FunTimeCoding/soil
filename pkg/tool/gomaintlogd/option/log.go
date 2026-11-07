package option

import "github.com/funtimecoding/soil/pkg/web/option"

type Log struct {
	option.Web
	PostgresLocator string
	LitePath        string
}
