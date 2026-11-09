package option

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Raid struct {
	option.Web
	PostgresLocator string
	LitePath        string
	LogCachePath    string
	ElitePath       string
	OutputPath      string
	Clock           func() time.Time
	Parser          func() face.Client
}
