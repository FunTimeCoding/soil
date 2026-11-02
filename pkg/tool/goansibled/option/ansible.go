package option

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"

type Ansible struct {
	Address         string
	ServiceTokens   []string
	Repository      string
	ClonePath       string
	AnsiblePath     string
	Playbook        []string
	PostgresLocator string
	LitePath        string
	Inventory       func() face.Client
}
