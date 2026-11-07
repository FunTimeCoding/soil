package option

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/web/option"
	"time"
)

type Salt struct {
	option.Web
	Metric          option.Web
	Repository      string
	ClonePath       string
	SaltPath        string
	PostgresLocator string
	LitePath        string
	KeySyncInterval time.Duration
	Master          func() face.SaltClient
}
