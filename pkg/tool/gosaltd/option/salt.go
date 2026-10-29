package option

import "github.com/funtimecoding/soil/pkg/provision/salt"

type Salt struct {
	Address         string
	ServiceTokens   []string
	Repository      string
	ClonePath       string
	SaltPath        string
	PostgresLocator string
	LitePath        string
	Master          func() *salt.Client
}
