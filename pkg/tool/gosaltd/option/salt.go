package option

import (
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"time"
)

type Salt struct {
	Address         string
	MetricAddress   string
	ServiceTokens   []string
	Repository      string
	ClonePath       string
	SaltPath        string
	PostgresLocator string
	LitePath        string
	KeySyncInterval time.Duration
	Master          func() face.SaltClient
}
