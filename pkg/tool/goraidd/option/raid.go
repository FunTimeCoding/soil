package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Raid struct {
	option.Web
	PostgresLocator string
	LitePath        string
	LogCachePath    string
	ElitePath       string
	OutputPath      string
	Parser          func() face.Client
}
