package option

import (
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"time"
)

type Salt struct {
	Address         string
	ServiceTokens   []string
	Repository      string
	ClonePath       string
	SaltPath        string
	PostgresLocator string
	LitePath        string
	KeySyncInterval time.Duration
	Master          func() *salt.Client
}
