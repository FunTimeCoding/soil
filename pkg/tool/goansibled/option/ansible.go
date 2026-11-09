package option

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/option"
)

type Ansible struct {
	option.Web
	Metric           option.Web
	Repository       string
	CloneDirectory   string
	CopyModes        bool
	AnsibleDirectory string
	Playbook         []string
	PostgresLocator  string
	LitePath         string
	Inventory        func() face.Client
}
