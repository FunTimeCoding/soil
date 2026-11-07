package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Ansible struct {
	option.Web
	Metric          option.Web
	Repository      string
	ClonePath       string
	AnsiblePath     string
	Playbook        []string
	PostgresLocator string
	LitePath        string
	Inventory       func() face.Client
}
