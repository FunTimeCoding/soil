package option

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	goclauded "github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Gitlab struct {
	option.Web
	Metric          option.Web
	Inventory       *instance.Inventory
	Client          func(*instance.Instance) (face.GitLabClient, error)
	LitePath        string
	PostgresLocator string
	Connector       func() goclauded.Client
	WatchInterval   time.Duration
	SweepInterval   time.Duration
}
